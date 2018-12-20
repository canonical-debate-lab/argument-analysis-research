package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/document"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/linker"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/linker/async"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/linker/storage"

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

	pool := linker.New(ctx, storage.NewLocalManager(ctx, "/db"), async.New)
	if err := pool.Load(ctx); err != nil {
		log.From(ctx).Fatal("loading pool", zap.Error(err))
	}

	srv.Router.Post("/argument/link", api.NewHandler(ctx, InsertHandler(ctx, pool, svc)))
	srv.Router.Get("/argument/links", api.NewHandler(ctx, ListHandler(ctx, pool, svc)))
	srv.Router.Get("/argument/linkers", api.NewHandler(ctx, LinkerHandler(ctx, pool, svc)))

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
func InsertHandler(ctx context.Context, pool *linker.Pool, svc Spec) api.HandlerFunc {
	type Request struct {
		ID        string               `json:"id"`
		Rater     string               `json:"rater"`
		Threshold float32              `json:"threshold"`
		Documents []*document.Document `json:"documents"`
	}

	type Response struct {
		*api.Error
		ID string `json:"id,omitempty"`
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

		var l *linker.Accessor
		if req.ID == "" {
			if req.Rater == "" {
				req.Rater = "https://research.democracy.ovh/argument/adw"
			}

			l, err = pool.Create(ctx, req.Rater, req.Threshold)
			if err != nil {
				resp.Fail(errors.Wrap(err, "creating new linker"))
				return resp
			}
		} else {
			l = pool.Get(ctx, req.ID)
			if l == nil || l.Linker == nil {
				resp.Fail(fmt.Errorf("no linker found for id: %s", req.ID))
				return resp
			}
		}

		resp.ID = l.ID
		ctx = log.WithFields(ctx, zap.String("linker", l.ID))

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
func ListHandler(ctx context.Context, pool *linker.Pool, svc Spec) api.HandlerFunc {
	type Response struct {
		Documents []*document.Document `json:"docs,omitempty"`
		Links     []*linker.Edge       `json:"links"`
		*api.Error
	}

	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) api.Response {
		var (
			resp = &Response{Error: &api.Error{}}
		)

		id := r.URL.Query().Get("id")
		if id == "" {
			resp.Fail(fmt.Errorf("linker id missing"))
			return resp
		}

		l := pool.Get(ctx, id)
		if l == nil || l.Linker == nil {
			resp.Fail(fmt.Errorf("no linker found for id: %s", r.URL.Query().Get("id")))
			return resp
		}

		noDocs := r.URL.Query().Get("docs")
		if noDocs != "false" {
			resp.Documents = l.ListDocuments(ctx)
			ctx = log.WithFields(ctx, zap.Int("documents", len(resp.Documents)))
		}

		resp.Links = l.ListLinks(ctx)

		return resp
	}
}

// LinkerHandler for getting a current list of links
func LinkerHandler(ctx context.Context, pool *linker.Pool, svc Spec) api.HandlerFunc {
	type Response struct {
		Linkers []linker.Linker `json:"linkers"`
		*api.Error
	}

	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) api.Response {
		var (
			resp = &Response{Error: &api.Error{}}
		)

		resp.Linkers = pool.List(ctx)

		return resp
	}
}
