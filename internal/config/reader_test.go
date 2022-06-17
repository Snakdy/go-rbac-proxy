package config_test

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/go-logr/logr/testr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.com/go-prism/go-rbac-proxy/internal/config"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	ctx := logr.NewContext(context.TODO(), testr.NewWithOptions(t, testr.Options{Verbosity: 10}))
	t.Run("normal config", func(t *testing.T) {
		c, err := config.Read(ctx, "./testdata/redis.yaml")
		assert.NoError(t, err)
		assert.ElementsMatch(t, c.Adapter.Redis.Addrs, []string{"localhost:6379"})
	})
	t.Run("embedded environment variables", func(t *testing.T) {
		require.NoError(t, os.Setenv("REDIS_PASSWORD", "hunter2"))
		c, err := config.Read(ctx, "./testdata/redis_env.yaml")
		assert.NoError(t, err)
		assert.EqualValues(t, "hunter2", c.Adapter.Redis.Password)
	})
}
