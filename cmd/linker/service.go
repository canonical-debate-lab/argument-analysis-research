package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/document"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/linker"

	"bitbucket.org/seibert-media/events/pkg/api"
	"bitbucket.org/seibert-media/events/pkg/service"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
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
	ctx := service.Init(svcKey, svcName, &svc)
	defer service.Defer(ctx)

	srv := api.New(":8080", svc.Debug)
	Routes(ctx, svc, srv)
	go srv.GracefulHandler(ctx)

	err := srv.Start(ctx)
	if err != nil {
		log.From(ctx).Fatal("running server", zap.Error(err))
	}

	log.From(ctx).Info("finished")
}

// Routes for this service
func Routes(ctx context.Context, svc Spec, srv *api.Server) {
	srv.Router.Route("/argument/segment", func(r chi.Router) {
		r.Post("/", api.NewHandler(ctx, Handler(ctx, svc)))
	})
}

// Handler for this endpoint
func Handler(ctx context.Context, svc Spec) api.HandlerFunc {
	type Response struct {
		Links []linker.DocumentLinks
		*api.Error
	}

	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) api.Response {
		var (
			req  []*document.Document
			resp = &Response{Error: &api.Error{}}
		)

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			resp.Fail(errors.Wrap(err, "decoding request"))
			return resp
		}

		ctx = log.WithFields(ctx,
			zap.Int("documents", len(req)),
		)

		httpRater := linker.NewHTTPRater("https://research.democracy.ovh/argument/adw")

		l := linker.New(httpRater, req, 0.5)
		links, err := l.Run(ctx)
		if err != nil {
			resp.Fail(errors.Wrap(err, "linking"))
		}

		resp.Links = links

		return resp
	}
}
