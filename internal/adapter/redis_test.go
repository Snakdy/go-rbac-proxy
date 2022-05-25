package adapter

import (
	"context"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"testing"
)

func newRedisAdapter(ctx context.Context, t *testing.T) *RedisAdapter {
	// set up the database
	rdb := miniredis.RunT(t)

	return NewRedisAdapter(ctx, &redis.UniversalOptions{
		Addrs: []string{rdb.Addr()},
	})
}
