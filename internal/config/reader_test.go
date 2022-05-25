package config_test

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/go-logr/logr/testr"
	"github.com/stretchr/testify/assert"
	"gitlab.com/go-prism/go-rbac-proxy/internal/config"
	"testing"
)

func TestRead(t *testing.T) {
	ctx := logr.NewContext(context.TODO(), testr.NewWithOptions(t, testr.Options{Verbosity: 10}))
	c, err := config.Read(ctx, "./testdata/redis.yaml")
	assert.NoError(t, err)
	assert.ElementsMatch(t, c.Adapter.Redis.Addrs, []string{"localhost:6379"})
}
