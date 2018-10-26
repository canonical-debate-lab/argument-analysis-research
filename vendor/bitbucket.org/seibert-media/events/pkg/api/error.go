package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/seibert-media/golibs/log"
	"go.uber.org/zap"
)

// Error implements the error interface and provides a matching implementation for Response
type Error struct {
	Err    error  `json:"-"`
	ErrStr string `json:"error,omitempty"`
}

// Fail sets the underlying error
func (e *Error) Fail(err error) {
	e.Err = err
	e.ErrStr = err.Error()
}

// Failure returns the underlying error
func (e *Error) Failure() error {
	return e.Err
}

// HTTPErr is a provisional convenience helper for handling errors inside the handler
// Setting usermsg to a non empty string, will overwrite the the error forwarded to the user with it's value
func HTTPErr(ctx context.Context, w http.ResponseWriter, err error) {
	log.From(ctx).Error("handling request", zap.Error(err))
	http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err), http.StatusInternalServerError)
}
