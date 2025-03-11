package service

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"
)

const DefaultHttpServerAddr = ":30000"

type ServerStartOpts struct {
	HttpServerAddr   string
	SocketServerAddr string
}

type Server struct {
	startOpts ServerStartOpts
	router    *chi.Mux
	ctx       context.Context
}

func NewServer(mux *chi.Mux, opts ServerStartOpts) *Server {
	return &Server{
		router:    mux,
		startOpts: opts,
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.startHttpServer()
	return nil
}

func (s *Server) startHttpServer() {
	// register the routes
	s.registerRoutes()

	// finally start the listener
	addr := DefaultHttpServerAddr
	if s.startOpts.HttpServerAddr != "" {
		addr = s.startOpts.HttpServerAddr
	}

	log.Info().Str("addr", addr).Msg("server started")
	go func() {
		http.ListenAndServe(addr, s.router)
	}()
}

func (s *Server) Stop() {
	log.Info().Msg("server stopped gracefully")
}

func (s *Server) registerRoutes() {

}
