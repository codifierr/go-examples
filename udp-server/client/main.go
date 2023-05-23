package main

import (
	"encoding/json"
	"net"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Message struct {
	TraceID string `json:"trace_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}

func main() {
	// Configure logging with zerolog
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Destination address for UDP connection
	destAddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"), // Replace with the IP address of the target UDP server
		Port: 64517,                    // Replace with the port number of the target UDP server
	}

	// Open a UDP connection
	conn, err := net.DialUDP("udp", nil, destAddr)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to dial UDP")
	}
	defer conn.Close()

	for {
		// Generate a trace ID using UUID
		traceID := uuid.New().String()

		// Create a sample JSON message with trace ID
		msg := Message{
			TraceID: traceID,
			Name:    "John Doe",
			Email:   "johndoe@example.com",
		}

		// Marshal the JSON message
		jsonData, err := json.Marshal(msg)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to marshal JSON")
		}

		// Send JSON data over UDP
		n, err := conn.Write(jsonData)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to write JSON data to UDP")
		}

		log.Info().Msgf("JSON data sent successfully over UDP. Sent %d bytes. Trace ID: %s", n, traceID)

		// Sleep for 1 second
		time.Sleep(1 * time.Second)
	}
}
