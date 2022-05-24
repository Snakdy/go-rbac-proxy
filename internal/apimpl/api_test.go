package apimpl

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/go-logr/logr/testr"
	"github.com/stretchr/testify/assert"
	"gitlab.com/go-prism/go-rbac-proxy/pkg/api"
	"testing"
)

var conf = Configuration{
	Globals: map[string][]string{
		"SUPER": {
			api.Verb_SUDO.String(),
		},
		"AUDIT": {
			api.Verb_READ.String(),
		},
	},
}

func TestAuthority_HasGlobalRole(t *testing.T) {
	ctx := logr.NewContext(context.TODO(), testr.NewWithOptions(t, testr.Options{
		Verbosity: 10,
	}))

	auth := NewAuthority(&conf, func(ctx context.Context, subject, role string) (bool, error) {
		if subject == "john.doe" {
			return true, nil
		}
		if subject == "jane.doe" && role == "AUDIT" {
			return true, nil
		}
		return false, nil
	}, nil)

	t.Run("admin can create", func(t *testing.T) {
		ok, err := auth.hasGlobalRole(ctx, &api.AccessRequest{
			Subject:  "john.doe",
			Resource: "foobar",
			Action:   api.Verb_CREATE,
		})
		assert.NoError(t, err)
		assert.True(t, ok)
	})
	t.Run("auditor cannot create", func(t *testing.T) {
		ok, err := auth.hasGlobalRole(ctx, &api.AccessRequest{
			Subject:  "jane.doe",
			Resource: "foobar",
			Action:   api.Verb_CREATE,
		})
		assert.NoError(t, err)
		assert.False(t, ok)
	})
}
