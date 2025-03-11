package main

import (
	"context"

	"github.com/bpatel8585/boring-sheet/pkg/service"
	"github.com/bpatel8585/boring-sheet/pkg/shutdown"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog/log"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	ctx := context.Background()

	server := service.NewServer(r, service.ServerStartOpts{})
	if err := server.Start(ctx); err != nil {
		// panic if failed to start the server
		log.Panic().Err(err).Msg("failed to start the server")
		return
	}

	<-shutdown.SignalChan()
	server.Stop()
}
