package adapter

import (
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/stretchr/testify/require"
	"math/rand/v2"
	"testing"
)

func newPostgres(t *testing.T) string {
	port := uint32(30000 + rand.IntN(2000))

	cfg := embeddedpostgres.DefaultConfig().
		Username("postgres").
		Password("hunter2").
		Database("something").
		Version(embeddedpostgres.V16).
		Port(port).
		BinariesPath(t.TempDir()).
		DataPath(t.TempDir()).
		RuntimePath(t.TempDir())

	postgres := embeddedpostgres.NewDatabase(cfg)
	require.NoError(t, postgres.Start())
	// cleanup when we're done
	t.Cleanup(func() {
		_ = postgres.Stop()
	})
	return cfg.GetConnectionURL()
}
