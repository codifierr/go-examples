package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/machinebox/graphql"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

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

	r := mux.NewRouter()
	for path, config := range configMap {
		r.HandleFunc(path, ProcessRequest).Methods(config.Method)
	}

	log.Info().Msg("Ready to server request at 8080")

	http.ListenAndServe(":8080", r)

}

func ProcessRequest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	template, _ := mux.CurrentRoute(r).GetPathTemplate()

	conf := configMap[template]

	vars := mux.Vars(r)

	var response []interface{}

	for _, backend := range conf.Backend {
		log.Info().Str("Endpoint", conf.Endpoint).Str("Backend", backend.Host[0]+backend.URLPattern).Msg("Serving Endpoint")
		graphQuery, err := getQuery(backend.Graphql.QueryPath)
		if err != nil {
			log.Error().AnErr("Error", err).Msg("Error in getting query")
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		// make a request
		req := graphql.NewRequest(graphQuery)

		//add query params
		for key, value := range r.URL.Query() {

			switch conf.InputQueryParams[key] {
			case "number":
				num, _ := strconv.Atoi(value[0])
				req.Var(key, num)
			default:
				req.Var(key, string(value[0]))
			}

		}
		// add path params
		for key, value := range vars {
			switch conf.InputPathParams[key] {
			case "number":
				num, _ := strconv.Atoi(value)
				req.Var(key, num)
			default:
				req.Var(key, string(value))
			}
		}

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

		// run it and capture the response
		var respBody interface{}
		if err := client.Run(r.Context(), req, &respBody); err != nil {
			log.Error().AnErr("Error", err).Msg("Error in getting response from graphql service")
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}
		response = append(response, respBody)
	}
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(response); err != nil {
		log.Printf("Can't encode for client %v - %s", response, err)
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

func getQuery(queryPath string) (string, error) {
	if queryPath != "" {
		q, err := os.ReadFile(queryPath)
		if err != nil {
			return "", err
		}

		return string(q), nil
	}
	return "", errors.New("query path is missing")
}

type Config struct {
	Endpoints []Endpoints `json:"endpoints"`
}

type Graphql struct {
	Type          string            `json:"type"`
	OperationName string            `json:"operationName"`
	Variables     map[string]string `json:"variables"`
	QueryPath     string            `json:"queryPath"`
}
type Backend struct {
	URLPattern string   `json:"urlPattern"`
	Method     string   `json:"method"`
	Graphql    Graphql  `json:"graphql"`
	Host       []string `json:"host"`
}
type Endpoints struct {
	Endpoint         string            `json:"endpoint"`
	Method           string            `json:"method"`
	InputPathParams  map[string]string `json:"inputPathParams,omitempty"`
	InputQueryParams map[string]string `json:"inputQueryParams,omitempty"`
	InputHeaders     []string          `json:"inputHeaders"`
	Backend          []Backend         `json:"backend"`
}
