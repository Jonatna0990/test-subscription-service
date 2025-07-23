package http

import (
	"fmt"
	"net/http"
)

type Server struct {
	srv    *http1.Server
	logger *logger.LogBuilder
}

func NewServer(cfg *server.Config, log *logger.LogBuilder) *Server {
	srv := http1.NewServer(cfg, log)
	return &Server{
		srv:    srv,
		logger: log,
	}
}

func (s *Server) RegisterRoute(route *Router) {

}

func (s *Server) Start() error {
	if s.srv.HasHandler() == false {
		return fmt.Errorf("no routes have registered")
	}

	err := s.srv.Run()
	if err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	return s.srv.Shutdown()
}
