package adapter

import (
	"context"
	"github.com/Snakdy/go-rbac-proxy/pkg/rbac"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newRedisAdapter(ctx context.Context, t *testing.T) *RedisAdapter {
	// set up the database
	rdb := miniredis.RunT(t)

	return NewRedisAdapter(ctx, &redis.UniversalOptions{
		Addrs: []string{rdb.Addr()},
	})
}

func TestRedisAdapter_ParseRoleBinding(t *testing.T) {
	p := rbac.Verb_SUDO
	var cases = []struct {
		in     string
		sub    string
		res    string
		action rbac.Verb
	}{
		{
			getKey("https://example.org/1", "foobar", &p),
			"https://example.org/1",
			"foobar",
			rbac.Verb_SUDO,
		},
		{
			getKey("https://example.org/1", "remote::alpine-aarnet", &p),
			"https://example.org/1",
			"remote::alpine-aarnet",
			rbac.Verb_SUDO,
		},
	}

	for _, tt := range cases {
		t.Run(tt.in, func(t *testing.T) {
			b := parseRoleBinding(tt.in)
			assert.EqualValues(t, tt.sub, b.GetSubject())
			assert.EqualValues(t, tt.res, b.GetResource())
			assert.EqualValues(t, tt.action, b.GetAction())
		})
	}
}
