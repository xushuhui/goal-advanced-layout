package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"goal-advanced-layout/pkg/log"
)

type Server struct {
	*gin.Engine
	httpSrv *http.Server
	host    string

	logger *log.Logger
}

type Option func(s *Server)

func NewServer(engine *gin.Engine, logger *log.Logger, opts ...Option) *Server {
	s := &Server{
		Engine: engine,
		logger: logger,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func WithServerHost(host string) Option {
	return func(s *Server) {
		s.host = host
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.httpSrv = &http.Server{
		Addr:    fmt.Sprintf("%s", s.host),
		Handler: s,
	}

	if err := s.httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.logger.Sugar().Fatalf("listen: %s\n", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Sugar().Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.httpSrv.Shutdown(ctx); err != nil {
		s.logger.Sugar().Fatal("Server forced to shutdown: ", err)
	}

	s.logger.Sugar().Info("Server exiting")
	return nil
}
