package service

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"contrib.go.opencensus.io/exporter/prometheus"
	"github.com/kelseyhightower/envconfig"
	"github.com/seibert-media/golibs/log"
	"go.opencensus.io/stats/view"
	"go.uber.org/zap"
)

// Spec is defining the common interface for service specification
type Spec interface {
	Base() BaseSpec
}

// BaseSpec stores all common service variables
type BaseSpec struct {
	Debug       bool   `default:"false" help:"set debug mode (default: false)"`
	SentryDSN   string `help:"sentry dsn key"`
	IsProd      bool   `envconfig:"is_prod" default:"false" help:"whether or not this service is running in production mode (used to disable certain features in dev)"`
	Local       bool   `envconfig:"local" default:"false" help:"setting local to true, enables human readable logging"`
	MetricsPort int    `envconfig:"metricsPort" default:"8080" help:"the metric server port (default: 8080)"`
}

// Base returns itself
func (b BaseSpec) Base() BaseSpec {
	return b
}

// GoogleSpec stores a ServiceAccount and an optional User field for Google Auth
type GoogleSpec struct {
	User           string `envconfig:"google_user" required:"false" help:"impersonation user for the google api"`
	ServiceAccount string `envconfig:"google_application_credentials" required:"true" help:"path to google service account file"`
}

// GoogleImpersonationSpec stores a ServiceAccount and a User field for Google Auth with impersonation
type GoogleImpersonationSpec struct {
	User           string `envconfig:"google_user" required:"true" help:"impersonation user for the google api"`
	ServiceAccount string `envconfig:"google_application_credentials" required:"true" help:"path to google service account file"`
}

// Init a service
// This contains our common boilerplate code required for every service inside the events project
func Init(key string, spec Spec) (ctx context.Context) {
	err := envconfig.Process("", spec)
	if err != nil {
		if perr, ok := err.(*envconfig.ParseError); ok && len(perr.Help) > 0 {
			fmt.Printf("ERROR: %s \n", perr.Err)
			fmt.Printf("HELP: %s: %s\n", perr.KeyName, perr.Help)
		} else {
			fmt.Println("failed parsing environment:", err)
		}

		os.Exit(1)
		return nil
	}

	runtime.GOMAXPROCS(runtime.NumCPU())

	logger, err := log.New(
		spec.Base().SentryDSN,
		spec.Base().Local,
	)
	if err != nil {
		panic(err)
	}

	if spec.Base().Debug {
		logger.SetLevel(zap.DebugLevel)
	}

	ctx = log.WithLogger(
		context.Background(),
		logger,
	)

	log.From(ctx).Info("preparing")

	// exporter, err := stackdriver.NewExporter(stackdriver.Options{
	// 	BundleDelayThreshold: time.Second / 10,
	// 	BundleCountThreshold: 10,
	// 	OnError: func(err error) {
	// 		log.From(ctx).Error("running stackdriver exporter", zap.Error(err))
	// 	},
	// 	DefaultTraceAttributes: map[string]interface{}{
	// 		"service": key,
	// 	},
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// trace.RegisterExporter(exporter)
	// trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	pe, err := prometheus.NewExporter(prometheus.Options{
		Namespace: "cdl",
	})
	if err != nil {
		panic(err)
	}
	view.RegisterExporter(pe)
	view.SetReportingPeriod(1 * time.Second)

	go metricsServer(ctx, spec.Base().MetricsPort, pe)

	return ctx
}

func metricsServer(ctx context.Context, port int, pe *prometheus.Exporter) {
	mux := http.NewServeMux()
	mux.Handle("/metrics", pe)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		log.From(ctx).Error("serving metrics", zap.Error(err))
		// TODO: don't break here -> rework/improve
		panic(err)
	}
}

// Defer contains all calls to be executed as a defer in main
func Defer(ctx context.Context) {
	log.From(ctx).Sync()
}
