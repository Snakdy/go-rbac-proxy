package apimpl

import (
	"github.com/Snakdy/go-rbac-proxy/internal/adapter"
	"github.com/Snakdy/go-rbac-proxy/internal/config"
	"github.com/Snakdy/go-rbac-proxy/pkg/rbac"
)

type Authority struct {
	rbac.UnimplementedAuthorityServer
	conf *config.Configuration

	receiver adapter.Adapter
}
