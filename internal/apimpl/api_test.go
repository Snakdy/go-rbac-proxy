package apimpl

import (
	"context"
	"github.com/Snakdy/go-rbac-proxy/internal/adapter"
	"github.com/Snakdy/go-rbac-proxy/internal/config"
	"github.com/Snakdy/go-rbac-proxy/pkg/rbac"
	"github.com/go-logr/logr"
	"github.com/go-logr/logr/testr"
	"github.com/stretchr/testify/assert"
	"testing"
)

var conf = config.Configuration{
	Globals: map[string][]string{
		"SUPER": {
			rbac.Verb_SUDO.String(),
		},
		"AUDIT": {
			rbac.Verb_READ.String(),
		},
	},
}

type testAdapter struct {
	adapter.Unimplemented
}

func (*testAdapter) SubjectHasGlobalRole(_ context.Context, subject, role string) (bool, error) {
	if subject == "john.doe" {
		return true, nil
	}
	if subject == "jane.doe" && role == "AUDIT" {
		return true, nil
	}
	return false, nil
}

func TestAuthority_HasGlobalRole(t *testing.T) {
	ctx := logr.NewContext(context.TODO(), testr.NewWithOptions(t, testr.Options{Verbosity: 10}))

	auth := NewAuthority(&conf, &testAdapter{})

	t.Run("admin can create", func(t *testing.T) {
		ok, err := auth.hasGlobalRole(ctx, &rbac.AccessRequest{
			Subject:  "john.doe",
			Resource: "foobar",
			Action:   rbac.Verb_CREATE,
		})
		assert.NoError(t, err)
		assert.True(t, ok)
	})
	t.Run("auditor cannot create", func(t *testing.T) {
		ok, err := auth.hasGlobalRole(ctx, &rbac.AccessRequest{
			Subject:  "jane.doe",
			Resource: "foobar",
			Action:   rbac.Verb_CREATE,
		})
		assert.NoError(t, err)
		assert.False(t, ok)
	})
}
