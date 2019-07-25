package main

import (
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/linker"
	linker_api "github.com/canonical-debate-lab/argument-analysis-research/pkg/linker/api"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/linker/async"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/linker/storage"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/api"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/service"

	"github.com/go-chi/chi"
	"github.com/seibert-media/golibs/log"
	"go.uber.org/zap"
)

const (
	svcName = "Argument Linker API Service"
	svcKey  = "argument-linker-api-service"
)

// Spec for the service
type Spec struct {
	service.BaseSpec
}

func main() {
	var svc Spec
	ctx := service.Init(svcKey, &svc)
	defer service.Defer(ctx)

	srv := api.New(":8080", svc.Debug)
	go srv.GracefulHandler(ctx)

	pool := linker.New(ctx, storage.NewLocalManager(ctx, "/db"), async.New)
	if err := pool.Load(ctx); err != nil {
		log.From(ctx).Fatal("loading pool", zap.Error(err))
	}

	srv.Router.Route("/argument", func(r chi.Router) {
		r.Post("/link", api.NewHandler(ctx, linker_api.InsertHandler(ctx, pool)))
		r.Get("/links", api.NewHandler(ctx, linker_api.ListHandler(ctx, pool)))
		r.Get("/linkers", api.NewHandler(ctx, linker_api.LinkerHandler(ctx, pool)))
	})

	err := srv.Start(ctx)
	if err != nil {
		log.From(ctx).Fatal("running server", zap.Error(err))
	}

	log.From(ctx).Info("finished")
}
