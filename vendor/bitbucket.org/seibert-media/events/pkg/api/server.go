package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/seibert-media/golibs/log"
	"go.uber.org/zap"

	"bitbucket.org/seibert-media/events/pkg/api/middlewares"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Server is our default wrapper for routing and serving
type Server struct {
	*http.Server

	Addr   string
	Router chi.Router
}

// New Server with default router and middlewares
func New(addr string, dbg bool) *Server {
	s := &Server{
		Addr:   addr,
		Router: chi.NewRouter(),
	}

	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middlewares.Logger)
	s.Router.Use(middleware.Timeout(60 * time.Second))

	if dbg {
		s.Router.Use(middlewares.DevMode)
	}

	s.Router.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf(`{"status": "ok"}`)))
	})

	return s
}

// Start the server
func (s *Server) Start(ctx context.Context) error {
	ctx = log.WithFields(ctx, zap.String("addr", s.Addr))
	s.Server = &http.Server{Addr: s.Addr, Handler: chi.ServerBaseContext(ctx, s.Router)}

	log.From(ctx).Info("serving")
	err := s.ListenAndServe()
	if err != http.ErrServerClosed {
		log.From(ctx).Error("serving", zap.Error(err))
		return s.Shutdown(ctx)
	}

	log.From(ctx).Info("stopping", zap.Error(err))
	return nil
}

// Shutdown the Server
func (s *Server) Shutdown(ctx context.Context) error {
	if s.Server != nil {
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		err := s.Server.Shutdown(ctx)
		if err != nil {
			log.From(ctx).Error("stopping gracefully", zap.Error(err))
			return err
		}
		s.Server = nil
	}

	return nil
}

// GracefulHandler watches for termination signals and takes care of graceful server shutdown
func (s *Server) GracefulHandler(ctx context.Context) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-c
	err := s.Shutdown(ctx)
	if err != nil {
		log.From(ctx).Fatal("handling signal", zap.Error(err))
	}
}
