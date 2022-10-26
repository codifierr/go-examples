package main

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var upgrader = websocket.Upgrader{}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error().AnErr("Error", err).Msg("Upgrade websocket")
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Error().AnErr("Error", err).Msg("Read websocket")
			break
		}
		log.Printf("recv: %s", message)
		log.Info().Bytes("Message", message).Msg("Received Message")
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Error().AnErr("Error", err).Msg("Write websocket")
			break
		}
	}
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Info().Msg("Initialized websocket echo server!")
	http.HandleFunc("/echo", echo)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal().AnErr("Error", err).Msg("Failed to Start Server")
	}
}
