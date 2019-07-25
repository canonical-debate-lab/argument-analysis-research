package api

import (
	"context"
	"net/http"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/api"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/linker"
)

// LinkerHandler for getting a current list of links
func LinkerHandler(ctx context.Context, pool *linker.Pool) api.HandlerFunc {
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
