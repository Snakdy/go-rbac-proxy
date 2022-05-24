package adapter

import (
	"context"
	"gitlab.com/go-prism/go-rbac-proxy/pkg/api"
)

type Adapter interface {
	SubjectHasGlobalRole(ctx context.Context, subject, role string) (bool, error)
	SubjectCanDoAction(ctx context.Context, subject, resource string, action api.Verb) (bool, error)

	Add(ctx context.Context, subject, resource string, action api.Verb) error
	AddGlobal(ctx context.Context, subject, role string) error
}

type SubjectHasGlobalRole = func(ctx context.Context, subject, role string) (bool, error)
type SubjectCanDoAction = func(ctx context.Context, subject, resource string, action api.Verb) (bool, error)
