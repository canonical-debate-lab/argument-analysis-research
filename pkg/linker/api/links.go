package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/api"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/document"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/linker"

	"github.com/seibert-media/golibs/log"
	"go.uber.org/zap"
)

// ListHandler for getting a current list of links
func ListHandler(ctx context.Context, pool *linker.Pool) api.HandlerFunc {
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
