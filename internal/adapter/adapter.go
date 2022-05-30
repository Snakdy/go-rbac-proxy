package adapter

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	"gitlab.com/go-prism/go-rbac-proxy/internal/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func New(ctx context.Context, c *config.Configuration) (Adapter, error) {
	ctx, span := otel.Tracer("").Start(ctx, "adapter_new", trace.WithAttributes(
		attribute.String("mode", c.Adapter.Mode),
	))
	defer span.End()
	log := logr.FromContextOrDiscard(ctx)
	log.V(1).Info("attempting to create adapter", "Mode", c.Adapter.Mode)
	switch c.Adapter.Mode {
	case "redis":
		return NewRedisAdapter(ctx, &c.Adapter.Redis), nil
	case "postgresql":
		fallthrough
	case "postgres":
		postgres, err := NewPostgresAdapter(ctx, c.Adapter.Postgres.DSN)
		if err != nil {
			return nil, err
		}
		return postgres, nil
	default:
		return nil, fmt.Errorf("unknown adapter: %s", c.Adapter.Mode)
	}
}
