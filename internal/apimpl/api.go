package apimpl

import (
	"context"
	"gitlab.com/go-prism/go-rbac-proxy/pkg/api"
)

func NewAuthority() *Authority {
	return &Authority{}
}

func (a *Authority) Can(ctx context.Context, request *api.AccessRequest) (*api.AccessResponse, error) {
	//TODO implement me
	panic("implement me")
}
