package apimpl

import (
	"gitlab.com/go-prism/go-rbac-proxy/internal/adapter"
	"gitlab.com/go-prism/go-rbac-proxy/internal/config"
	"gitlab.com/go-prism/go-rbac-proxy/pkg/rbac"
)

type Authority struct {
	rbac.UnimplementedAuthorityServer
	conf *config.Configuration

	receiver adapter.Adapter
}
