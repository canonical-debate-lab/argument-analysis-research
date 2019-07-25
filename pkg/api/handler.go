package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/seibert-media/golibs/log"
	"go.uber.org/zap"
)

// HandlerFunc represents a http.HandlerFunc returning an error
type HandlerFunc func(context.Context, http.ResponseWriter, *http.Request) Response

// Response with basic error handling features
type Response interface {
	// Fail should set the internal error value if existing
	Fail(error)
	// Failure should return the internal error value if existing (or nil)
	Failure() error
}

// NewHandler returns a standard http.HandlerFunc wrapping our default encoding and error handling
func NewHandler(ctx context.Context, f HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := f(r.Context(), w, r)

		if err := resp.Failure(); err != nil {
			log.From(ctx).Error("handling api", zap.Error(err))
		}

		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			HTTPErr(ctx, w, errors.Wrap(err, "encoding response"))
		}
	}
}
