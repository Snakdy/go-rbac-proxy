package adapter

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/go-logr/logr/testr"
	"github.com/stretchr/testify/assert"
	"gitlab.com/go-prism/go-rbac-proxy/pkg/api"
	"testing"
)

// interface guard
var _ Adapter = &PostgresAdapter{}

func TestPostgresAdapter_Add(t *testing.T) {
	ctx := logr.NewContext(context.TODO(), testr.NewWithOptions(t, testr.Options{Verbosity: 10}))
	// set up the database
	postgres := newPostgres(t)
	defer postgres.Stop()

	adapter, err := NewPostgresAdapter(ctx, "user=prism password=hunter2 dbname=prism host=localhost port=5432 sslmode=disable")
	assert.NoError(t, err)

	// create duplicate roles
	assert.NoError(t, adapter.Add(ctx, "john.doe", "foo", api.Verb_SUDO))
	assert.NoError(t, adapter.Add(ctx, "jane.doe", "foo", api.Verb_SUDO))
	assert.NoError(t, adapter.Add(ctx, "john.doe", "foo", api.Verb_SUDO))
}

func TestNewPostgresAdapter(t *testing.T) {
	ctx := logr.NewContext(context.TODO(), testr.NewWithOptions(t, testr.Options{Verbosity: 10}))
	// set up the database
	postgres := newPostgres(t)
	defer postgres.Stop()

	adapter, err := NewPostgresAdapter(ctx, "user=prism password=hunter2 dbname=prism host=localhost port=5432 sslmode=disable")
	assert.NoError(t, err)

	// create some roles
	assert.NoError(t, adapter.Add(ctx, "john.doe", "foo", api.Verb_SUDO))
	assert.NoError(t, adapter.Add(ctx, "jane.doe", "foo", api.Verb_READ))

	// verify the roles
	var cases = []struct {
		name string
		sub  string
		res  string
		verb api.Verb
		ok   bool
	}{
		{
			"sudo can create",
			"john.doe",
			"foo",
			api.Verb_CREATE,
			true,
		},
		{
			"reader can read",
			"jane.doe",
			"foo",
			api.Verb_READ,
			true,
		},
		{
			"reader cannot delete",
			"jane.doe",
			"foo",
			api.Verb_DELETE,
			false,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ok, err := adapter.SubjectCanDoAction(ctx, tt.sub, tt.res, tt.verb)
			assert.NoError(t, err)
			assert.EqualValues(t, tt.ok, ok)
		})
	}
}
