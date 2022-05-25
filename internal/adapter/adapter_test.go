package adapter

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/go-logr/logr/testr"
	"github.com/stretchr/testify/assert"
	"gitlab.com/go-prism/go-rbac-proxy/pkg/api"
	"testing"
)

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
			assert.NoError(t, tt.adapter.Add(ctx, "john.doe", "foo", api.Verb_SUDO))
			assert.NoError(t, tt.adapter.Add(ctx, "jane.doe", "foo", api.Verb_SUDO))
			assert.NoError(t, tt.adapter.Add(ctx, "john.doe", "foo", api.Verb_SUDO))
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
			assert.NoError(t, tt.adapter.Add(ctx, "john.doe", "foo", api.Verb_SUDO))
			assert.NoError(t, tt.adapter.Add(ctx, "jane.doe", "foo", api.Verb_READ))

			// verify the roles
			var subCases = []struct {
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
