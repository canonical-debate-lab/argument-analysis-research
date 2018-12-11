package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/document"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/linker"

	"bitbucket.org/seibert-media/events/pkg/api"
	"bitbucket.org/seibert-media/events/pkg/service"
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

	linker := linker.New(linker.NewHTTPRater("https://research.democracy.ovh/argument/adw"), 0.45)
	go linker.Run(ctx)

	srv.Router.Post("/argument/link", api.NewHandler(ctx, InsertHandler(ctx, linker, svc)))
	srv.Router.Get("/argument/links", api.NewHandler(ctx, ListHandler(ctx, linker, svc)))

	err := srv.Start(ctx)
	if err != nil {
		log.From(ctx).Fatal("running server", zap.Error(err))
	}

	log.From(ctx).Info("finished")
}

// Routes for this service
func Routes(ctx context.Context, svc Spec, srv *api.Server) {
}

// InsertHandler for adding documents to the linker
func InsertHandler(ctx context.Context, l *linker.Linker, svc Spec) api.HandlerFunc {
	type Request struct {
		Documents []*document.Document `json:"documents"`
	}

	type Response struct {
		*api.Error
	}

	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) api.Response {
		var (
			req  Request
			resp = &Response{Error: &api.Error{}}
		)

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			resp.Fail(errors.Wrap(err, "decoding request"))
			return resp
		}

		ctx = log.WithFields(ctx,
			zap.Int("documents", len(req.Documents)),
		)

		for _, doc := range req.Documents {
			err := l.InsertDocument(ctx, doc)
			if err != nil {
				resp.Fail(errors.Wrap(err, "inserting document"))
				break
			}
		}

		return resp
	}
}

// ListHandler for getting a current list of links
func ListHandler(ctx context.Context, l *linker.Linker, svc Spec) api.HandlerFunc {
	type Response struct {
		Documents []*document.Document `json:"docs,omitempty"`
		Links     []*linker.Edge       `json:"links"`
		*api.Error
	}

	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) api.Response {
		var (
			resp = &Response{Error: &api.Error{}}
		)

		noDocs := r.URL.Query().Get("docs")
		if noDocs != "false" {
			resp.Documents = l.ListDocuments(ctx)
			ctx = log.WithFields(ctx, zap.Int("documents", len(resp.Documents)))
		}

		resp.Links = l.ListLinks(ctx)

		return resp
	}
}
