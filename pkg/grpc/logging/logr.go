package grpc_logr

import (
	"context"
	"github.com/go-logr/logr"
	"google.golang.org/grpc"
)

type LogrInterceptor struct {
	log logr.Logger
}

func NewLogrInterceptor(log logr.Logger) *LogrInterceptor {
	return &LogrInterceptor{
		log: log,
	}
}

func (i *LogrInterceptor) Unary(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	log := i.log.WithValues("Method", info.FullMethod, "Server", info.Server)
	log.V(2).Info("processing gRPC request")
	return handler(logr.NewContext(ctx, log), req)
}
