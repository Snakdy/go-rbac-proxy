package apimpl

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/go-logr/logr/testr"
	"github.com/stretchr/testify/assert"
	"gitlab.com/go-prism/go-rbac-proxy/internal/config"
	"gitlab.com/go-prism/go-rbac-proxy/pkg/rbac"
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

func TestAuthority_HasGlobalRole(t *testing.T) {
	ctx := logr.NewContext(context.TODO(), testr.NewWithOptions(t, testr.Options{Verbosity: 10}))

	auth := NewAuthority(&conf, func(ctx context.Context, subject, role string) (bool, error) {
		if subject == "john.doe" {
			return true, nil
		}
		if subject == "jane.doe" && role == "AUDIT" {
			return true, nil
		}
		return false, nil
	}, nil, nil, nil)

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
