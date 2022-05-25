package adapter

import (
	"context"
	"errors"
	"github.com/djcass44/go-utils/orm"
	"github.com/go-logr/logr"
	"gitlab.com/go-prism/go-rbac-proxy/pkg/api"
	"gitlab.com/go-prism/go-rbac-proxy/pkg/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type PostgresAdapter struct {
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
	log := logr.FromContextOrDiscard(ctx).WithValues("Subject", subject, "Role", role).WithName("postgres")
	log.V(1).Info("checking if subject has global role")
	var count int64
	if err := p.db.Model(&schemas.PostgresRoleBinding{}).Where("subject = ? AND resource = ?", subject, role).Count(&count).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.V(1).Info("matching role could not be found")
			return false, nil
		}
		log.Error(err, "failed to check for global role membership")
		return false, err
	}
	return count > 0, nil
}

func (p *PostgresAdapter) SubjectCanDoAction(ctx context.Context, subject, resource string, action api.Verb) (bool, error) {
	log := logr.FromContextOrDiscard(ctx).WithValues("Subject", subject, "Resource", resource, "Action", action.String()).WithName("postgres")
	log.V(1).Info("checking if subject has role")
	var count int64
	if err := p.db.Model(&schemas.PostgresRoleBinding{}).Where("subject = ? AND resource = ? AND (verb = ? OR verb = 'SUDO')", subject, resource, action.String()).Count(&count).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.V(1).Info("matching role could not be found")
			return false, nil
		}
		log.Error(err, "failed to check for role membership")
		return false, err
	}
	return count > 0, nil
}

func (p *PostgresAdapter) Add(ctx context.Context, subject, resource string, action api.Verb) error {
	log := logr.FromContextOrDiscard(ctx).WithValues("Subject", subject, "Resource", resource, "Action", action.String()).WithName("postgres")
	log.V(1).Info("creating role binding")
	if err := p.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&schemas.PostgresRoleBinding{
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
	log := logr.FromContextOrDiscard(ctx).WithValues("Subject", subject, "Role").WithName("postgres")
	log.V(1).Info("creating global role binding")
	if err := p.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&schemas.PostgresRoleBinding{
		Subject:  subject,
		Resource: role,
	}).Error; err != nil {
		log.Error(err, "failed to add role binding")
		return err
	}
	return nil
}
