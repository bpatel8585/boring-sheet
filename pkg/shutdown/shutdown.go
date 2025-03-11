package shutdown

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

const timeout = 10 * time.Second

var signalChannel chan os.Signal
var shutdownChannel chan os.Signal

func init() {
	shutdownChannel = make(chan os.Signal)
	signalChannel = make(chan os.Signal)
	go shutdownHandler()
	signal.Notify(shutdownChannel, syscall.SIGTERM)
	signal.Notify(shutdownChannel, syscall.SIGINT)
}

// SignalChan is used as a global channel for signaling shutdowns
func SignalChan() chan os.Signal {
	return signalChannel
}

func shutdownHandler() {
	signal := <-shutdownChannel
	log.Info().Str("signal", signal.String()).Msg("sending shutdown signals")

	// Notifies listeners of shutdown
	close(signalChannel)

	// wait for timeout
	time.Sleep(timeout)
	log.Info().Msg("shutting down ...")
	os.Exit(0)
}
