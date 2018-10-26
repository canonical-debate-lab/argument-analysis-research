package middlewares

import (
	"net/http"
	"time"

	"github.com/seibert-media/golibs/log"
	"go.uber.org/zap"
)

// Logger the request to the logger in its own context
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		start := time.Now()
		sw := statusWriter{ResponseWriter: w}
		next.ServeHTTP(&sw, r)
		duration := time.Now().Sub(start)

		log.From(ctx).Info("handling",
			zap.Int("status", sw.status), zap.String("method", r.Method),
			zap.String("host", r.Host), zap.String("path", r.RequestURI),
			zap.Duration("duration", duration),
		)
	})
}
