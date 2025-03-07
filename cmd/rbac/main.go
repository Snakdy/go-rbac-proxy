package main

import (
	"context"
	"github.com/Snakdy/go-rbac-proxy/internal/adapter"
	"github.com/Snakdy/go-rbac-proxy/internal/apimpl"
	"github.com/Snakdy/go-rbac-proxy/internal/config"
	grpclogr "github.com/Snakdy/go-rbac-proxy/pkg/grpc/logging"
	"github.com/Snakdy/go-rbac-proxy/pkg/rbac"
	"github.com/djcass44/go-utils/logging"
	"github.com/djcass44/go-utils/otel"
	"github.com/djcass44/go-utils/otel/metrics"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.com/autokubeops/serverless"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
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
	Otel       struct {
		Enabled     bool    `split_words:"true"`
		Environment string  `envconfig:"GITLAB_ENVIRONMENT_NAME"`
		SampleRate  float64 `split_words:"true" default:"0.05"`
	}
}

const ServiceName = "go-rbac-proxy"

func main() {
	// read environment
	var e environment
	envconfig.MustProcess("app", &e)

	// configure logging
	zc := zap.NewProductionConfig()
	zc.Level = zap.NewAtomicLevelAt(zapcore.Level(e.LogLevel * -1))
	log, ctx := logging.NewZap(context.TODO(), zc)

	// setup otel
	err := otel.Build(ctx, otel.Options{
		Enabled:       e.Otel.Enabled,
		ServiceName:   ServiceName,
		Environment:   e.Otel.Environment,
		KubeNamespace: os.Getenv("KUBE_NAMESPACE"),
		SampleRate:    e.Otel.SampleRate,
	})
	if err != nil {
		log.Error(err, "failed to setup tracing")
		os.Exit(1)
		return
	}

	metrics.MustNewDefault(ctx)

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
	logRpc := grpclogr.NewLogrInterceptor(log)
	// create and configure the server
	gsrv := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.UnaryInterceptor(logRpc.Unary),
	)
	rbac.RegisterAuthorityServer(gsrv, apimpl.NewAuthority(c, adp))

	// configure routing
	router := mux.NewRouter()
	router.Use(logging.Middleware(log))
	router.Handle("/metrics", promhttp.Handler()).Methods(http.MethodGet)
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
