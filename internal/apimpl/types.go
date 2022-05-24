package apimpl

import (
	"context"
	"gitlab.com/go-prism/go-rbac-proxy/pkg/api"
)

type Globals map[string][]string

type Authority struct {
	api.UnimplementedAuthorityServer
	conf *Configuration

	subjectHasGlobalRole SubjectHasGlobalRole
	subjectCanDoAction   SubjectCanDoAction
}

type Configuration struct {
	Globals Globals
}

type SubjectHasGlobalRole = func(ctx context.Context, subject, role string) (bool, error)
type SubjectCanDoAction = func(ctx context.Context, subject, resource string, action api.Verb) (bool, error)
