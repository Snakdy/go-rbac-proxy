package adapter

import (
	"context"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-logr/logr"
	"github.com/go-logr/logr/testr"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"gitlab.com/go-prism/go-rbac-proxy/internal/config"
	"gitlab.com/go-prism/go-rbac-proxy/pkg/rbac"
	"testing"
)

func TestNew(t *testing.T) {
	ctx := logr.NewContext(context.TODO(), testr.NewWithOptions(t, testr.Options{Verbosity: 10}))

	// test that postgres is configured correctly
	t.Run("postgres", func(t *testing.T) {
		postgres := newPostgres(t)
		defer postgres.Stop()

		adp, err := New(ctx, &config.Configuration{
			Adapter: config.Adapter{
				Mode: "postgres",
				Postgres: config.PostgresAdapter{
					DSN: "user=prism password=hunter2 dbname=prism host=localhost port=5432 sslmode=disable",
				},
			},
		})
		assert.NoError(t, err)
		assert.NotNil(t, adp)
	})

	// test that redis is configured correctly
	t.Run("redis", func(t *testing.T) {
		rdb := miniredis.RunT(t)

		adp, err := New(ctx, &config.Configuration{
			Adapter: config.Adapter{
				Mode: "redis",
				Redis: redis.UniversalOptions{
					Addrs: []string{rdb.Addr()},
				},
			},
		})
		assert.NoError(t, err)
		assert.NotNil(t, adp)
	})

	// test that unknown adapters err
	t.Run("unknown adapter errors", func(t *testing.T) {
		adp, err := New(ctx, &config.Configuration{
			Adapter: config.Adapter{
				Mode: "foobar",
			},
		})
		assert.Error(t, err)
		assert.Nil(t, adp)
	})
}

func TestAdapter_Add(t *testing.T) {
	ctx := logr.NewContext(context.TODO(), testr.NewWithOptions(t, testr.Options{Verbosity: 10}))

	var cases = []struct {
		name    string
		adapter Adapter
	}{
		{
			"postgresql",
			newPostgresAdapter(ctx, t),
		},
		{
			"redis",
			newRedisAdapter(ctx, t),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			// create duplicate roles
			assert.NoError(t, tt.adapter.Add(ctx, "john.doe", "foo", rbac.Verb_SUDO))
			assert.NoError(t, tt.adapter.Add(ctx, "jane.doe", "foo", rbac.Verb_SUDO))
			assert.NoError(t, tt.adapter.Add(ctx, "john.doe", "foo", rbac.Verb_SUDO))
		})
	}
}

func TestNewAdapter(t *testing.T) {
	ctx := logr.NewContext(context.TODO(), testr.NewWithOptions(t, testr.Options{Verbosity: 10}))

	var cases = []struct {
		name    string
		adapter Adapter
	}{
		{
			"postgresql",
			newPostgresAdapter(ctx, t),
		},
		{
			"redis",
			newRedisAdapter(ctx, t),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			// create some roles
			assert.NoError(t, tt.adapter.Add(ctx, "john.doe", "foo", rbac.Verb_SUDO))
			assert.NoError(t, tt.adapter.Add(ctx, "jane.doe", "foo", rbac.Verb_READ))

			// verify the roles
			var subCases = []struct {
				name string
				sub  string
				res  string
				verb rbac.Verb
				ok   bool
			}{
				{
					"sudo can create",
					"john.doe",
					"foo",
					rbac.Verb_CREATE,
					true,
				},
				{
					"reader can read",
					"jane.doe",
					"foo",
					rbac.Verb_READ,
					true,
				},
				{
					"reader cannot delete",
					"jane.doe",
					"foo",
					rbac.Verb_DELETE,
					false,
				},
			}
			for _, tts := range subCases {
				t.Run(tts.name, func(t *testing.T) {
					ok, err := tt.adapter.SubjectCanDoAction(ctx, tts.sub, tts.res, tts.verb)
					assert.NoError(t, err)
					assert.EqualValues(t, tts.ok, ok)
				})
			}
		})
	}
}

func TestAdapter_List(t *testing.T) {
	ctx := logr.NewContext(context.TODO(), testr.NewWithOptions(t, testr.Options{Verbosity: 10}))

	var cases = []struct {
		name    string
		adapter Adapter
	}{
		{
			"postgresql",
			newPostgresAdapter(ctx, t),
		},
		{
			"redis",
			newRedisAdapter(ctx, t),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			// create data
			assert.NoError(t, tt.adapter.Add(ctx, "john.doe", "foo", rbac.Verb_SUDO))
			assert.NoError(t, tt.adapter.Add(ctx, "jane.doe", "foo", rbac.Verb_SUDO))
			assert.NoError(t, tt.adapter.AddGlobal(ctx, "jane.doe", "FOOBAR"))
			assert.NoError(t, tt.adapter.AddGlobal(ctx, "john.doe", "FOOBAR"))

			// list by subject
			t.Run("list by subject", func(t *testing.T) {
				res, err := tt.adapter.ListBySub(ctx, "john.doe")
				assert.NoError(t, err)
				assert.Len(t, res, 2)
			})

			// list by role/resource
			t.Run("list by role or resource", func(t *testing.T) {
				res, err := tt.adapter.ListByRole(ctx, "FOOBAR")
				assert.NoError(t, err)
				assert.Len(t, res, 2)
			})

			// list
			t.Run("list all", func(t *testing.T) {
				res, err := tt.adapter.List(ctx)
				assert.NoError(t, err)
				assert.Len(t, res, 4)
			})
		})
	}
}
