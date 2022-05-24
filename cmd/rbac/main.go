package main

import (
	"context"
	"github.com/djcass44/go-utils/logging"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"gitlab.com/autokubeops/serverless"
	"go.uber.org/zap"
	"net/http"
)

type environment struct {
	Port int `envconfig:"PORT" default:"8080"`
}

func main() {
	// read environment
	var e environment
	envconfig.MustProcess("app", &e)

	// configure logging
	zc := zap.NewProductionConfig()
	log, _ := logging.NewZap(context.TODO(), zc)

	// configure routing
	router := mux.NewRouter()
	router.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("OK"))
	})

	// start serving
	serverless.NewBuilder(router).
		WithLogger(log).
		WithPort(e.Port).
		Run()
}
