package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rsocket/rsocket-go"
	"github.com/rsocket/rsocket-go/payload"
	"github.com/rsocket/rsocket-go/rx/mono"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Caller().Logger()

	logLevelStr := os.Getenv("LOG_LEVEL")
	logLevel, err := zerolog.ParseLevel(logLevelStr)
	if err != nil {
		log.Error().Err(err).Msg("Error in setting log level defaulting to info")
		logLevel = zerolog.InfoLevel
	}

	// set the global log level for zerolog
	zerolog.SetGlobalLevel(logLevel)

	log.Info().Msg("Initializing rsocket server!")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	address := os.Getenv("ADDRESS")
	if address == "" {
		address = ":7878"
	}

	errChan := make(chan error)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			cancel()
		}
	}()

	go func() {
		err = rsocket.Receive().
			Acceptor(func(ctx context.Context, setup payload.SetupPayload, sendingSocket rsocket.CloseableRSocket) (rsocket.RSocket, error) {
				// bind responder
				return rsocket.NewAbstractSocket(
					rsocket.RequestResponse(func(msg payload.Payload) mono.Mono {
						log.Debug().Bytes("Message", msg.Data()).Msg("Message received")
						return mono.Just(msg)
					}),
				), nil
			}).
			Transport(rsocket.WebsocketServer().SetAddr(address).Build()).Serve(ctx)
		log.Error().AnErr("Error", err).Msg("Failed to start server")
		errChan <- err
	}()

	go func() {
		err := <-errChan
		log.Error().Err(err).Msg("Fatal error. main will exit now!")
		cancel()
	}()
	<-ctx.Done()
}
