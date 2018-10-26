package service

import (
	"context"
	"fmt"
	"os"
	"runtime"

	"bitbucket.org/seibert-media/events/pkg/version"

	"github.com/kelseyhightower/envconfig"
	"github.com/seibert-media/golibs/log"
)

// Spec is defining the common interface for service specification
type Spec interface {
	Base() BaseSpec
}

// BaseSpec stores all common service variables
type BaseSpec struct {
	Debug       bool   `default:"false" help:"set debug mode (default: false)"`
	SentryDSN   string `help:"sentry dsn key"`
	ShowVersion bool   `default:"true" help:"show version info"`
	ProjectID   string `envconfig:"project" required:"true" help:"project id for pubsub"`
	IsProd      bool   `envconfig:"is_prod" default:"false" help:"whether or not this service is running in production mode (used to disable certain features in dev)"`
	Local       bool   `envconfig:"local" default:"false" help:"setting local to true, enables pretty logging"`
}

// Base returns itself
func (b BaseSpec) Base() BaseSpec {
	return b
}

// GoogleSpec stores a ServiceAccount and a User field for Google Auth
type GoogleSpec struct {
	User           string `envconfig:"google_user" required:"true" help:"impersonation user for the reseller api"`
	ServiceAccount string `envconfig:"google_application_credentials" required:"true" help:"path to google service account file"`
}

// Init a service
// This contains our common boilerplate code required for every service inside the events project
func Init(key, name string, spec Spec) (ctx context.Context) {
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
	version.Print(spec.Base().ShowVersion, name)

	logger, err := log.New(
		spec.Base().SentryDSN,
		spec.Base().Debug,
		spec.Base().Local,
	)
	if err != nil {
		panic(err)
	}
	logger = logger.
		WithRelease(version.Release()).
		WithFields(
			version.Fields(spec.Base().Debug, spec.Base().Local, key)...,
		)

	ctx = log.WithLogger(
		context.Background(),
		logger,
	)

	log.From(ctx).Info("preparing")

	return ctx
}

// Defer contains all calls to be executed as a defer in main
func Defer(ctx context.Context) {
	log.From(ctx).Sync()
}
