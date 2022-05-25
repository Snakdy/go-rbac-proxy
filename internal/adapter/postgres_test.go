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
	postgres := newPostgres(t)
	t.Cleanup(func() {
		_ = postgres.Stop()
	})

	adapter, err := NewPostgresAdapter(ctx, "user=prism password=hunter2 dbname=prism host=localhost port=5432 sslmode=disable")
	assert.NoError(t, err)

	return adapter
}
