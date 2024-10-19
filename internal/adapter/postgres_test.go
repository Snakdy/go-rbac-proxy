package adapter

import (
	"context"
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/stretchr/testify/assert"
	"testing"
)

// interface guard
var _ Adapter = &PostgresAdapter{}

func newPostgresAdapter(ctx context.Context, t *testing.T, version embeddedpostgres.PostgresVersion) *PostgresAdapter {
	// set up the database
	dsn := newPostgres(t, version)

	adapter, err := NewPostgresAdapter(ctx, dsn)
	assert.NoError(t, err)

	return adapter
}
