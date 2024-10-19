package adapter

import (
	"context"
	"github.com/Snakdy/go-rbac-proxy/pkg/rbac"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Unimplemented struct{}

func (*Unimplemented) SubjectHasGlobalRole(context.Context, string, string) (bool, error) {
	return false, status.Error(codes.Unimplemented, "unimplemented")
}

func (*Unimplemented) SubjectCanDoAction(context.Context, string, string, rbac.Verb) (bool, error) {
	return false, status.Error(codes.Unimplemented, "unimplemented")
}

func (*Unimplemented) Add(context.Context, string, string, rbac.Verb) error {
	return status.Error(codes.Unimplemented, "unimplemented")
}

func (*Unimplemented) AddGlobal(context.Context, string, string) error {
	return status.Error(codes.Unimplemented, "unimplemented")
}

func (*Unimplemented) ListBySub(context.Context, string) ([]*rbac.RoleBinding, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func (*Unimplemented) ListByRole(context.Context, string) ([]*rbac.RoleBinding, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func (*Unimplemented) List(context.Context) ([]*rbac.RoleBinding, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
