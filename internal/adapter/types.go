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
}

type SubjectHasGlobalRole = func(ctx context.Context, subject, role string) (bool, error)
type SubjectCanDoAction = func(ctx context.Context, subject, resource string, action rbac.Verb) (bool, error)
type Add = func(ctx context.Context, subject, resource string, action rbac.Verb) error
type AddGlobal = func(ctx context.Context, subject, role string) error
