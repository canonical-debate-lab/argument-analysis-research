package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/api"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/document"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/linker"

	"github.com/pkg/errors"
	"github.com/seibert-media/golibs/log"
	"go.uber.org/zap"
)

// InsertHandler for adding documents to the linker
func InsertHandler(ctx context.Context, pool *linker.Pool) api.HandlerFunc {
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
