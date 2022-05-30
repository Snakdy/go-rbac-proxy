package adapter

import (
	"context"
	"errors"
	"github.com/djcass44/go-utils/orm"
	"github.com/go-logr/logr"
	otelgorm "github.com/kostyay/gorm-opentelemetry"
	"gitlab.com/go-prism/go-rbac-proxy/pkg/rbac"
	"gitlab.com/go-prism/go-rbac-proxy/pkg/schemas"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type PostgresAdapter struct {
	Unimplemented
	db *gorm.DB
}

func NewPostgresAdapter(ctx context.Context, dsn string) (*PostgresAdapter, error) {
	log := logr.FromContextOrDiscard(ctx).WithName("postgres")
	log.V(2).Info("attempting to connect to database", "DSN", dsn)
	// open connection
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: orm.NewGormLogger(log, time.Millisecond*200),
	})
	if err != nil {
		log.Error(err, "failed to open database connection")
		return nil, err
	}
	log.Info("established database connection")

	log.V(1).Info("enabling plugins")
	if err := database.Use(otelgorm.NewPlugin()); err != nil {
		log.Error(err, "failed to enable SQL OpenTelemetry plugin")
	}

	// migrate
	log.V(1).Info("running database migrations")
	if err := database.AutoMigrate(&schemas.PostgresRoleBinding{}); err != nil {
		log.Error(err, "failed to run auto-migration")
		return nil, err
	}
	log.V(1).Info("completed database migrations")

	return &PostgresAdapter{
		db: database,
	}, nil
}

func (p *PostgresAdapter) SubjectHasGlobalRole(ctx context.Context, subject, role string) (bool, error) {
	ctx, span := otel.Tracer("").Start(ctx, "adapter_postgres_subjectHasGlobalRole", trace.WithAttributes(
		attribute.String("subject", subject),
		attribute.String("role", role),
	))
	defer span.End()
	log := logr.FromContextOrDiscard(ctx).WithValues("Subject", subject, "Role", role).WithName("postgres")
	log.V(1).Info("checking if subject has global role")
	var count int64
	if err := p.db.WithContext(ctx).Model(&schemas.PostgresRoleBinding{}).Where("subject = ? AND resource = ?", subject, role).Count(&count).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.V(1).Info("matching role could not be found")
			return false, nil
		}
		log.Error(err, "failed to check for global role membership")
		return false, err
	}
	return count > 0, nil
}

func (p *PostgresAdapter) SubjectCanDoAction(ctx context.Context, subject, resource string, action rbac.Verb) (bool, error) {
	ctx, span := otel.Tracer("").Start(ctx, "adapter_postgres_subjectCanDoAction", trace.WithAttributes(
		attribute.String("subject", subject),
		attribute.String("resource", resource),
		attribute.String("action", action.String()),
	))
	defer span.End()
	log := logr.FromContextOrDiscard(ctx).WithValues("Subject", subject, "Resource", resource, "Action", action.String()).WithName("postgres")
	log.V(1).Info("checking if subject has role")
	var count int64
	if err := p.db.WithContext(ctx).Model(&schemas.PostgresRoleBinding{}).Where("subject = ? AND resource = ? AND (verb = ? OR verb = 'SUDO')", subject, resource, action.String()).Count(&count).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.V(1).Info("matching role could not be found")
			return false, nil
		}
		log.Error(err, "failed to check for role membership")
		return false, err
	}
	return count > 0, nil
}

func (p *PostgresAdapter) Add(ctx context.Context, subject, resource string, action rbac.Verb) error {
	ctx, span := otel.Tracer("").Start(ctx, "adapter_postgres_add", trace.WithAttributes(
		attribute.String("subject", subject),
		attribute.String("resource", resource),
		attribute.String("action", action.String()),
	))
	defer span.End()
	log := logr.FromContextOrDiscard(ctx).WithValues("Subject", subject, "Resource", resource, "Action", action.String()).WithName("postgres")
	log.V(1).Info("creating role binding")
	if err := p.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&schemas.PostgresRoleBinding{
		Subject:  subject,
		Resource: resource,
		Verb:     action.String(),
	}).Error; err != nil {
		log.Error(err, "failed to add role binding")
		return err
	}
	return nil
}

func (p *PostgresAdapter) AddGlobal(ctx context.Context, subject, role string) error {
	ctx, span := otel.Tracer("").Start(ctx, "adapter_postgres_addGlobal", trace.WithAttributes(
		attribute.String("subject", subject),
		attribute.String("role", role),
	))
	defer span.End()
	log := logr.FromContextOrDiscard(ctx).WithValues("Subject", subject, "Role").WithName("postgres")
	log.V(1).Info("creating global role binding")
	if err := p.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&schemas.PostgresRoleBinding{
		Subject:  subject,
		Resource: role,
	}).Error; err != nil {
		log.Error(err, "failed to add role binding")
		return err
	}
	return nil
}

func (p *PostgresAdapter) ListBySub(ctx context.Context, subject string) ([]*rbac.RoleBinding, error) {
	ctx, span := otel.Tracer("").Start(ctx, "adapter_postgres_listBySub", trace.WithAttributes(
		attribute.String("subject", subject),
	))
	defer span.End()
	log := logr.FromContextOrDiscard(ctx).WithValues("Subject", subject).WithName("postgres")
	log.V(1).Info("scanning for roles with subject")

	var results []schemas.PostgresRoleBinding
	if err := p.db.WithContext(ctx).Where("subject = ?", subject).Find(&results).Error; err != nil {
		log.Error(err, "failed to find role bindings by subject")
		return nil, err
	}

	items := make([]*rbac.RoleBinding, len(results))
	for i := range results {
		items[i] = &rbac.RoleBinding{
			Subject:  results[i].Subject,
			Resource: results[i].Resource,
			Action:   rbac.Verb(rbac.Verb_value[results[i].Verb]),
		}
	}
	log.Info("successfully fetched roles for subject", "Count", len(items))
	return items, nil
}

func (p *PostgresAdapter) ListByRole(ctx context.Context, role string) ([]*rbac.RoleBinding, error) {
	ctx, span := otel.Tracer("").Start(ctx, "adapter_postgres_listByRole", trace.WithAttributes(
		attribute.String("role", role),
	))
	defer span.End()
	log := logr.FromContextOrDiscard(ctx).WithValues("Role", role).WithName("postgres")
	log.V(1).Info("scanning for roles with subject")

	var results []schemas.PostgresRoleBinding
	if err := p.db.WithContext(ctx).Where("resource = ?", role).Find(&results).Error; err != nil {
		log.Error(err, "failed to find role bindings by role")
		return nil, err
	}

	items := make([]*rbac.RoleBinding, len(results))
	for i := range results {
		items[i] = &rbac.RoleBinding{
			Subject:  results[i].Subject,
			Resource: results[i].Resource,
			Action:   rbac.Verb(rbac.Verb_value[results[i].Verb]),
		}
	}
	log.Info("successfully fetched bindings for role", "Count", len(items))
	return items, nil
}

func (p *PostgresAdapter) List(ctx context.Context) ([]*rbac.RoleBinding, error) {
	ctx, span := otel.Tracer("").Start(ctx, "adapter_postgres_list")
	defer span.End()
	log := logr.FromContextOrDiscard(ctx).WithName("postgres")
	log.V(1).Info("scanning for roles")

	var results []schemas.PostgresRoleBinding
	if err := p.db.WithContext(ctx).Find(&results).Error; err != nil {
		log.Error(err, "failed to find role bindings")
		return nil, err
	}

	items := make([]*rbac.RoleBinding, len(results))
	for i := range results {
		items[i] = &rbac.RoleBinding{
			Subject:  results[i].Subject,
			Resource: results[i].Resource,
			Action:   rbac.Verb(rbac.Verb_value[results[i].Verb]),
		}
	}
	log.Info("successfully fetched bindings", "Count", len(items))
	return items, nil
}
