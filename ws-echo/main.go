package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var upgrader = websocket.Upgrader{}

type StreamMessage struct {
	Message    string `json:"message"`
	ReplyCount int    `json:"replyCount"`
	Delay      int    `json:"delay"`
}

type Stream struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

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
		log.Info().Bytes("Message", message).Msg("Received Message")
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Error().AnErr("Error", err).Msg("Write websocket")
			break
		}
	}
}

func echostream(w http.ResponseWriter, r *http.Request) {
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
		log.Info().Bytes("Message", message).Msg("Received Message")
		streamMessage := &StreamMessage{
			Delay:      1,
			ReplyCount: 10,
		}
		err = json.Unmarshal(message, &streamMessage)
		if err != nil {
			log.Error().AnErr("Error", err).Msg("Invalid message")
		}

		log.Info().Interface("Stream Message", streamMessage).Msg("Stream Message")

		data := streamMessage.Message
		count := streamMessage.ReplyCount
		duration := time.Duration(streamMessage.Delay) * time.Second
		for {
			if count > 0 {
				stream := Stream{
					Message: data,
					Id:      uuid.New().String(),
				}
				streamBytes, err := json.Marshal(stream)
				if err != nil {
					log.Error().AnErr("Error", err).Msg("Marshal error")
					break
				}
				log.Info().Interface("Stream Message", streamMessage).Msg("Writing Message")
				err = c.WriteMessage(mt, streamBytes)
				if err != nil {
					log.Error().AnErr("Error", err).Msg("Write websocket")
					break
				}
			} else {
				break
			}
			log.Info().Int("Duration", int(duration)).Msg("Sleeping")
			time.Sleep(duration)
			count--
		}
	}
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Info().Msg("Initialized websocket echo server!")
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/echostream", echostream)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal().AnErr("Error", err).Msg("Failed to Start Server")
	}
}
