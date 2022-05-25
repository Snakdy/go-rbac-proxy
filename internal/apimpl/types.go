package apimpl

import (
	"gitlab.com/go-prism/go-rbac-proxy/internal/adapter"
	"gitlab.com/go-prism/go-rbac-proxy/internal/config"
	"gitlab.com/go-prism/go-rbac-proxy/pkg/api"
)

type Authority struct {
	api.UnimplementedAuthorityServer
	conf *config.Configuration

	subjectHasGlobalRole adapter.SubjectHasGlobalRole
	subjectCanDoAction   adapter.SubjectCanDoAction

	addRole       adapter.Add
	addGlobalRole adapter.AddGlobal
}
