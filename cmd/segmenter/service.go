package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/document"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/api"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/service"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"github.com/seibert-media/golibs/log"
	"go.uber.org/zap"
)

const (
	svcName = "Argument Segmentation API Service"
	svcKey  = "argument-segmentation-api-service"
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
		r.Post("/bulk", api.NewHandler(ctx, BulkHandler(ctx, svc)))
	})
}

// Handler for this endpoint
func Handler(ctx context.Context, svc Spec) api.HandlerFunc {
	type Request struct {
		Input string `json:"input"`
	}

	type Response struct {
		*document.Document
		*api.Error
	}

	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) api.Response {
		var (
			req  *Request
			resp = &Response{Error: &api.Error{}}
		)

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			resp.Fail(errors.Wrap(err, "decoding request"))
			return resp
		}

		ctx = log.WithFields(ctx,
			zap.String("input", req.Input),
		)

		doc, err := document.New(ctx, req.Input)
		if err != nil {
			resp.Fail(errors.Wrap(err, "analyzing input"))
			return resp
		}

		resp.Document = doc

		return resp
	}
}

// BulkHandler for this endpoint
func BulkHandler(ctx context.Context, svc Spec) api.HandlerFunc {
	type Request struct {
		Input []string `json:"inputs"`
	}

	type Response struct {
		Documents []*document.Document `json:"documents"`
		*api.Error
	}

	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) api.Response {
		var (
			req  *Request
			resp = &Response{Error: &api.Error{}}
		)

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			resp.Fail(errors.Wrap(err, "decoding request"))
			return resp
		}

		for _, str := range req.Input {

			ctx := log.WithFields(ctx,
				zap.String("input", str),
			)

			doc, err := document.New(ctx, str)
			if err != nil {
				resp.Fail(errors.Wrap(err, "analyzing input"))
				return resp
			}

			resp.Documents = append(resp.Documents, doc)
		}

		return resp
	}
}
