package apimpl

import (
	"context"
	"github.com/djcass44/go-utils/utilities/sliceutils"
	"github.com/go-logr/logr"
	"gitlab.com/go-prism/go-rbac-proxy/internal/adapter"
	"gitlab.com/go-prism/go-rbac-proxy/internal/config"
	"gitlab.com/go-prism/go-rbac-proxy/pkg/rbac"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewAuthority(conf *config.Configuration, receiver adapter.Adapter) *Authority {
	return &Authority{
		conf:     conf,
		receiver: receiver,
	}
}

func (a *Authority) Can(ctx context.Context, request *rbac.AccessRequest) (*rbac.GenericResponse, error) {
	ctx, span := otel.Tracer("").Start(ctx, "api_authority_can", trace.WithAttributes(
		attribute.String("subject", request.GetSubject()),
		attribute.String("resource", request.GetResource()),
		attribute.String("action", request.GetAction().String()),
	))
	defer span.End()
	log := logr.FromContextOrDiscard(ctx).WithValues("Subject", request.GetSubject(), "Resource", request.GetResource(), "Action", request.GetAction().String())
	log.Info("checking if subject can perform action on resource")
	// check global roles
	ok, err := a.hasGlobalRole(ctx, request)
	if err != nil {
		return nil, err
	}
	if ok {
		metricCan.Add(ctx, 1, metric.WithAttributes(attribute.String(attributeSourceKey, sourceGlobal)))
		return &rbac.GenericResponse{Message: "", Ok: true}, nil
	}
	ok, err = a.hasRole(ctx, request)
	if ok {
		metricCan.Add(ctx, 1, metric.WithAttributes(attribute.String(attributeSourceKey, sourceRole)))
		return &rbac.GenericResponse{Message: "", Ok: true}, nil
	}
	metricCan.Add(ctx, 1, metric.WithAttributes(attribute.String(attributeSourceKey, sourceNone)))
	return &rbac.GenericResponse{Message: "", Ok: false}, nil
}

func (a *Authority) AddRole(ctx context.Context, request *rbac.AddRoleRequest) (*rbac.GenericResponse, error) {
	attributes := []attribute.KeyValue{
		attribute.String("subject", request.GetSubject()),
		attribute.String("resource", request.GetResource()),
		attribute.String("action", request.GetAction().String()),
	}
	ctx, span := otel.Tracer("").Start(ctx, "api_authority_addRole", trace.WithAttributes(attributes...))
	defer span.End()
	log := logr.FromContextOrDiscard(ctx).WithValues("Subject", request.GetSubject(), "Resource", request.GetResource(), "Action", request.GetAction().String())
	log.Info("creating role binding to role")
	metricAdd.Add(ctx, 1, metric.WithAttributes(attributes...))
	if err := a.receiver.Add(ctx, request.GetSubject(), request.GetResource(), request.GetAction()); err != nil {
		return nil, err
	}
	return &rbac.GenericResponse{
		Message: "",
		Ok:      true,
	}, nil
}

func (a *Authority) ListBySub(ctx context.Context, request *rbac.ListBySubRequest) (*rbac.ListResponse, error) {
	ctx, span := otel.Tracer("").Start(ctx, "api_authority_listBySub", trace.WithAttributes(
		attribute.String("subject", request.GetSubject()),
	))
	defer span.End()
	log := logr.FromContextOrDiscard(ctx).WithValues("Subject", request.GetSubject())
	log.Info("listing role bindings for subject")
	items, err := a.receiver.ListBySub(ctx, request.GetSubject())
	if err != nil {
		return nil, err
	}
	return &rbac.ListResponse{Results: items}, nil
}

func (a *Authority) ListByRole(ctx context.Context, request *rbac.ListByRoleRequest) (*rbac.ListResponse, error) {
	ctx, span := otel.Tracer("").Start(ctx, "api_authority_listByRole", trace.WithAttributes(
		attribute.String("role", request.GetRole()),
	))
	defer span.End()
	log := logr.FromContextOrDiscard(ctx).WithValues("Role", request.GetRole())
	log.Info("listing role bindings for role")
	items, err := a.receiver.ListByRole(ctx, request.GetRole())
	if err != nil {
		return nil, err
	}
	return &rbac.ListResponse{Results: items}, nil
}

func (a *Authority) List(ctx context.Context, _ *emptypb.Empty) (*rbac.ListResponse, error) {
	ctx, span := otel.Tracer("").Start(ctx, "api_authority_list")
	defer span.End()
	log := logr.FromContextOrDiscard(ctx)
	log.Info("listing role bindings")
	items, err := a.receiver.List(ctx)
	if err != nil {
		return nil, err
	}
	return &rbac.ListResponse{Results: items}, nil
}

func (a *Authority) AddGlobalRole(ctx context.Context, request *rbac.AddGlobalRoleRequest) (*rbac.GenericResponse, error) {
	attributes := []attribute.KeyValue{
		attribute.String("subject", request.GetSubject()),
		attribute.String("role", request.GetRole()),
	}
	ctx, span := otel.Tracer("").Start(ctx, "api_authority_addGlobalRole", trace.WithAttributes(attributes...))
	defer span.End()
	log := logr.FromContextOrDiscard(ctx).WithValues("Subject", request.GetSubject(), "Role", request.GetRole())
	log.Info("creating role binding to global role")
	metricAddGlobal.Add(ctx, 1, metric.WithAttributes(attributes...))
	if err := a.receiver.AddGlobal(ctx, request.GetSubject(), request.GetRole()); err != nil {
		return nil, err
	}
	return &rbac.GenericResponse{
		Message: "",
		Ok:      true,
	}, nil
}

func canPerformAction(actions []string, verb rbac.Verb) bool {
	// if the role has no actions, skip it
	if len(actions) == 0 {
		return false
	}
	// if the only action is SUDO don't skip it
	if len(actions) == 1 {
		if actions[0] == rbac.Verb_SUDO.String() {
			return true
		}
	}
	// otherwise check that the requested action
	// is in the list of supported actions
	return sliceutils.Includes(actions, verb.String())
}

func (a *Authority) hasRole(ctx context.Context, request *rbac.AccessRequest) (bool, error) {
	ctx, span := otel.Tracer("").Start(ctx, "api_authority_hasRole", trace.WithAttributes(
		attribute.String("subject", request.GetSubject()),
		attribute.String("resource", request.GetResource()),
		attribute.String("action", request.GetAction().String()),
	))
	defer span.End()
	log := logr.FromContextOrDiscard(ctx).WithValues("Subject", request.GetSubject(), "Resource", request.GetResource(), "Action", request.GetAction().String(), "Global", false)
	log.V(1).Info("checking roles")

	ok, err := a.receiver.SubjectCanDoAction(ctx, request.GetSubject(), request.GetResource(), request.GetAction())
	if err != nil {
		return false, err
	}
	log.Info("successfully checked role membership", "Member", ok)
	return ok, nil
}

func (a *Authority) hasGlobalRole(ctx context.Context, request *rbac.AccessRequest) (bool, error) {
	ctx, span := otel.Tracer("").Start(ctx, "api_authority_hasGlobalRole", trace.WithAttributes(
		attribute.String("subject", request.GetSubject()),
		attribute.String("resource", request.GetResource()),
		attribute.String("action", request.GetAction().String()),
	))
	defer span.End()
	log := logr.FromContextOrDiscard(ctx).WithValues("Subject", request.GetSubject(), "Resource", request.GetResource(), "Action", request.GetAction().String(), "Global", true)
	log.V(1).Info("checking global roles")
	// iterate through each of the global
	// roles
	for k, v := range a.conf.Globals {
		log = log.WithValues("Role", k)
		log.V(2).Info("checking global role")
		// check if the role matches the
		// verb that we're requesting
		if !canPerformAction(v, request.GetAction()) {
			log.V(2).Info("skipping global role as it does not include our Verb")
			continue
		}
		log.V(2).Info("verb matched, checking if subject has the role")
		// check if the subject actually has
		// the role
		ok, err := a.receiver.SubjectHasGlobalRole(ctx, request.GetSubject(), k)
		if err != nil {
			log.Error(err, "failed to check if subject has role")
			return false, err
		}
		log.Info("successfully checked role membership", "Member", ok)
		if ok {
			return true, nil
		}
	}
	log.Info("successfully checked role membership", "Member", false)
	return false, nil
}
