package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/machinebox/graphql"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var upgrader = websocket.Upgrader{}

var (
	configMap map[string]Endpoints
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Info().Msg("Initializing app...")

	configPath := "config.json"

	configMap = make(map[string]Endpoints)

	err := loadAndStoreConfig(configPath, configMap)
	if err != nil {
		log.Panic().AnErr("Error", err).Msg("Exiting")
	}

	log.Info().Msg("Config Load complete")

	for path, config := range configMap {
		log.Info().Str("Method", config.Method).Msg("HTTP Method to be used")
		if config.Method == "STREAM" {
			http.HandleFunc(path, ProcessStream)
		}
	}

	log.Info().Msg("Ready to server request at 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal().AnErr("Error", err).Msg("Failed to Start Server")
	}
}

func ProcessStream(w http.ResponseWriter, r *http.Request) {
	template := r.URL.Path
	log.Debug().Str("Template", template).Msg("Path to pick config")

	conf := configMap[template]

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error().AnErr("Error", err).Msg("Upgrade websocket")
		return
	}
	defer c.Close()

	payloadChan := make(chan Payload)
	responseChan := make(chan interface{})
	cleanup := make(chan bool)

	go func() {
		for {
			select {
			case <-cleanup:
				return
			default:
				_, message, err := c.ReadMessage()
				if err != nil {
					log.Error().AnErr("Error", err).Msg("Read websocket")
					cleanup <- true
					break
				}
				log.Info().Bytes("Message", message).Msg("Received Message")
				var payload Payload
				err = json.Unmarshal(message, &payload)
				if err != nil {
					log.Error().AnErr("Error", err).Msg("Invalid message")
				}
				payloadChan <- payload
			}
		}
	}()
	go func() {
		for {
			select {
			case <-cleanup:
				return
			case res := <-responseChan:
				data, err := json.Marshal(res)
				if err != nil {
					log.Printf("Can't marshal  %v - %s", res, err)
				}
				err = c.WriteMessage(1, data)
				if err != nil {
					log.Error().AnErr("Error", err).Msg("Write websocket")
					cleanup <- true
					return
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case <-cleanup:
				return
			case payload := <-payloadChan:
				log.Info().Interface("Payload", payload)
				for _, backend := range conf.Backend {
					log.Info().Str("Endpoint", conf.Endpoint).Str("Backend", backend.Host[0]+backend.URLPattern).Msg("Serving Endpoint")

					// make a request
					req := graphql.NewRequest(payload.Query)

					// Todo: Create a single client for all backend and share across request
					// create a client (safe to share across requests)
					client := graphql.NewClient(backend.Host[0] + backend.URLPattern)

					//Pass requested headers to backend
					for _, header := range conf.InputHeaders {
						if val := r.Header.Get(header); val != "" {
							log.Debug().Str("HeaderName", header).Str("Value", val).Msg("Passed Header")
							req.Header.Set(header, val)
						}
					}
					req.Header.Set("Content-Type", "application/json")
					// add variables to the request from payload
					addVariablesFromBody(req, &payload)

					// run it and capture the response
					var respBody interface{}
					if err := client.Run(r.Context(), req, &respBody); err != nil {
						log.Error().AnErr("Error", err).Msg("Error in getting response from graphql service")
						http.Error(w, "Internal Error", http.StatusInternalServerError)
						return
					}
					responseChan <- respBody
				}
			}
		}
	}()
	<-cleanup
	log.Info().Msg("Request terminated either by client or in case of unusable websocket")
}

func addVariablesFromBody(req *graphql.Request, vars *Payload) {
	for k, v := range vars.Vars {
		req.Var(k, v)
	}
}

func loadAndStoreConfig(configPath string, configMap map[string]Endpoints) error {
	if configPath != "" {
		file, err := os.ReadFile(configPath)
		if err != nil {
			return err
		}
		var config Config
		_ = json.Unmarshal([]byte(file), &config)
		for _, v := range config.Endpoints {
			configMap[v.Endpoint] = v
		}
	}
	return nil
}

type Config struct {
	Endpoints []Endpoints `json:"endpoints"`
}

type Backend struct {
	URLPattern string   `json:"urlPattern"`
	Method     string   `json:"method"`
	Host       []string `json:"host"`
}
type Endpoints struct {
	Endpoint     string    `json:"endpoint"`
	Method       string    `json:"method"`
	InputHeaders []string  `json:"inputHeaders"`
	Backend      []Backend `json:"backend"`
}
type Payload struct {
	Query string                 `json:"query"`
	Vars  map[string]interface{} `json:"vars"`
}
