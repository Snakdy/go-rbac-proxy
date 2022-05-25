package adapter

import (
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/stretchr/testify/require"
	"testing"
)

func newPostgres(t *testing.T) *embeddedpostgres.EmbeddedPostgres {
	postgres := embeddedpostgres.NewDatabase(embeddedpostgres.DefaultConfig().
		//BinaryRepositoryURL("https://prism.v2.dcas.dev/api/v1/maven/-").
		Username("prism").
		Password("hunter2").
		Database("prism").
		Version(embeddedpostgres.V14).
		BinariesPath(t.TempDir()).
		DataPath(t.TempDir()).
		RuntimePath(t.TempDir()),
	)
	require.NoError(t, postgres.Start())
	return postgres
}
