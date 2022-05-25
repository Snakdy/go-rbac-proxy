package apimpl

import (
	"gitlab.com/go-prism/go-rbac-proxy/internal/adapter"
	"gitlab.com/go-prism/go-rbac-proxy/pkg/api"
)

type Globals map[string][]string

type Authority struct {
	api.UnimplementedAuthorityServer
	conf *Configuration

	subjectHasGlobalRole adapter.SubjectHasGlobalRole
	subjectCanDoAction   adapter.SubjectCanDoAction

	addRole       adapter.Add
	addGlobalRole adapter.AddGlobal
}

type Configuration struct {
	Globals Globals
}
