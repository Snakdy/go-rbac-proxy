package adapter

import (
	"context"
	"gitlab.com/go-prism/go-rbac-proxy/pkg/rbac"
)

type Adapter interface {
	SubjectHasGlobalRole(ctx context.Context, subject, role string) (bool, error)
	SubjectCanDoAction(ctx context.Context, subject, resource string, action rbac.Verb) (bool, error)

	Add(ctx context.Context, subject, resource string, action rbac.Verb) error
	AddGlobal(ctx context.Context, subject, role string) error

	ListBySub(ctx context.Context, subject string) ([]*rbac.RoleBinding, error)
	ListByRole(ctx context.Context, role string) ([]*rbac.RoleBinding, error)
}
