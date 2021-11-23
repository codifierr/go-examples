package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting http client")
	counter := 1
	for {
		log.Info().Int("Counter", counter).Msg("Counter")
		makeHttpCall()
		makePostCall()
		time.Sleep(5 * time.Second)
		counter++
		if counter == 10 {
			log.Info().Msg("Stopping http client")
			break
		}
	}
}

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

func makePostCall() {
	resp, err := http.Post("https://httpbin.org/post", "application/json", nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to send request message")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read response body")
	}

	log.Info().Str("Post Response", string(body)).Msg("Response")
}
