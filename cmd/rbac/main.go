package main

import (
	"context"
	"github.com/djcass44/go-utils/logging"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"gitlab.com/autokubeops/serverless"
	"gitlab.com/go-prism/go-rbac-proxy/internal/apimpl"
	"gitlab.com/go-prism/go-rbac-proxy/pkg/api"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"net/http"
)

type environment struct {
	Port     int `envconfig:"PORT" default:"8080"`
	LogLevel int `split_words:"true"`
}

func main() {
	// read environment
	var e environment
	envconfig.MustProcess("app", &e)

	// configure logging
	zc := zap.NewProductionConfig()
	zc.Level = zap.NewAtomicLevelAt(zapcore.Level(e.LogLevel * -1))
	log, _ := logging.NewZap(context.TODO(), zc)

	// configure grpc
	gsrv := grpc.NewServer()
	api.RegisterAuthorityServer(gsrv, apimpl.NewAuthority(nil, nil, nil, nil, nil))

	// configure routing
	router := mux.NewRouter()
	router.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("OK"))
	})

	// start serving
	serverless.NewBuilder(router).
		WithLogger(log).
		WithPort(e.Port).
		WithGRPC(gsrv).
		Run()
}
