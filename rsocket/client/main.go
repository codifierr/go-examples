package main

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/rsocket/rsocket-go"
	"github.com/rsocket/rsocket-go/payload"
)

func main() {
	// Connect to server
	cli, err := rsocket.Connect().
		SetupPayload(payload.NewString("Hello", "World")).
		Transport(rsocket.WebsocketClient().SetURL("ws://127.0.0.1:31539").Build()).
		Start(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect")
	}
	defer cli.Close()
	// Send request
	result, err := cli.RequestResponse(payload.NewString("Hi, this is satyendra", "Hello")).Block(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to send request")
	}
	log.Info().Bytes("Response", result.Data()).Msg("Response")
}
