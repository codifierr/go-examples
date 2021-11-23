package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

// main is the entry point for the wasm module
func main() {
	log.Info().Msg("Starting http client")
	// counter to stop the program after 10 calls
	counter := 1
	for {
		log.Info().Int("Counter", counter).Msg("Counter")
		makeHttpCall()

		// sleep for 5 second before making another call
		time.Sleep(5 * time.Second)

		counter++
		if counter == 10 {
			log.Info().Msg("Stopping http client")
			break
		}
	}
}

// makeHttpCall makes a http call to the server
func makeHttpCall() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to send request message")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read response body")
	}
	log.Info().Str("Get Response", string(body)).Msg("Response")
}
