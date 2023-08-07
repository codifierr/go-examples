package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	nrgin "github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var router = gin.Default()

func Ping(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/plain")
	c.Writer.WriteString("Pong!")
}

func Hello(c *gin.Context) {
	name := c.Params.ByName("name")
	c.Writer.Header().Set("Content-Type", "text/plain")
	c.Writer.WriteString("Hello, \n" + name)
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	logLevelStr := os.Getenv("LOG_LEVEL")
	logLevel, err := zerolog.ParseLevel(logLevelStr)
	if err != nil {
		log.Error().Err(err).Msg("Error in setting log level defaulting to info")
		logLevel = zerolog.InfoLevel
	}

	// set the global log level for zerolog
	zerolog.SetGlobalLevel(logLevel)
	newRelicKey := os.Getenv("NEWRELIC_API_KEY")
	newRelicEnabled, err := strconv.ParseBool(os.Getenv("NEWRELIC_ENABLED"))
	if err != nil || newRelicKey == "" {
		newRelicEnabled = false
	}

	nr, _ := newrelic.NewApplication(
		newrelic.ConfigAppName("Http-Ping"),
		newrelic.ConfigLicense(newRelicKey),
		newrelic.ConfigDebugLogger(os.Stdout),
		newrelic.ConfigEnabled(newRelicEnabled),
		newrelic.ConfigDistributedTracerEnabled(true),
	)

	router.Use(nrgin.Middleware(nr))

	router.GET("/ping", Ping)
	router.GET("/hello/:name", Hello)

	log.Info().Msg("Initialized Ping Server!")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal().AnErr("Error", err).Msg("Failed to Start Server")
	}
}
