package main

import (
	"net"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Configure logging with zerolog
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Info().Msg("Starting UDP server...")

	// listen to incoming udp packets
	udpServer, err := net.ListenPacket("udp", ":8080")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start UDP server")
	}
	defer udpServer.Close()

	buf := make([]byte, 32768) // buffer for reading packets
	for {
		n, addr, err := udpServer.ReadFrom(buf)
		if err != nil {
			log.Error().Err(err).Msg("Failed to read from UDP connection")
			continue
		}
		go response(addr, buf[:n])
	}
}

func response(addr net.Addr, buf []byte) {
	log.Info().Str("Message", string(buf)).Interface("Address", addr).Msg("Message recieved")
}
