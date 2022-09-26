package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Pong!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")

	fmt.Fprint(w, "Hello, \n"+name)
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	router := httprouter.New()
	router.GET("/ping", Ping)
	router.GET("/hello/:name", Hello)

	log.Info().Msg("Initialized Ping Server!")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal().AnErr("Error", err).Msg("Failed to Start Server")
	}
}
