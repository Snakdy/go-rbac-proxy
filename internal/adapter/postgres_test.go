package adapter

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

// interface guard
var _ Adapter = &PostgresAdapter{}

func newPostgresAdapter(ctx context.Context, t *testing.T) *PostgresAdapter {
	// set up the database
	dsn := newPostgres(t)

	adapter, err := NewPostgresAdapter(ctx, dsn)
	assert.NoError(t, err)

	return adapter
}
