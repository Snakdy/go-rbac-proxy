package main

import (
	"context"
	"github.com/djcass44/go-utils/logging"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"gitlab.com/autokubeops/serverless"
	"gitlab.com/go-prism/go-rbac-proxy/internal/adapter"
	"gitlab.com/go-prism/go-rbac-proxy/internal/apimpl"
	"gitlab.com/go-prism/go-rbac-proxy/internal/config"
	grpc_logr "gitlab.com/go-prism/go-rbac-proxy/pkg/grpc/logging"
	"gitlab.com/go-prism/go-rbac-proxy/pkg/rbac"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"os"
)

type environment struct {
	Port       int    `envconfig:"PORT" default:"8080"`
	LogLevel   int    `split_words:"true"`
	ConfigPath string `split_words:"true" required:"true"`
}

func main() {
	// read environment
	var e environment
	envconfig.MustProcess("app", &e)

	// configure logging
	zc := zap.NewProductionConfig()
	zc.Level = zap.NewAtomicLevelAt(zapcore.Level(e.LogLevel * -1))
	log, ctx := logging.NewZap(context.TODO(), zc)

	c, err := config.Read(ctx, e.ConfigPath)
	if err != nil {
		os.Exit(1)
		return
	}
	adp, err := adapter.New(ctx, c)
	if err != nil {
		log.Error(err, "failed to instantiate adapter")
		os.Exit(1)
		return
	}

	// configure grpc
	logRpc := grpc_logr.NewLogrInterceptor(log)
	// create and configure the server
	gsrv := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(logRpc.Unary),
	)
	rbac.RegisterAuthorityServer(gsrv, apimpl.NewAuthority(c, adp.SubjectHasGlobalRole, adp.SubjectCanDoAction, adp.Add, adp.AddGlobal))

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
