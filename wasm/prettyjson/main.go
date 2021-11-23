package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/rs/zerolog/log"
)

// main is the entry point for the wasm module
func main() {
	log.Info().Msg("Go wasm")
	js.Global().Set("formatJSON", jsonWrapper())
	<-make(chan bool)
}

// prettyJson takes a json string and returns a pretty printed json string
func prettyJson(input string) (string, error) {
	var raw interface{}
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

// jsonWrapper is a wrapper function for prettyJson
func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		inputJSON := args[0].String()
		log.Info().Str("Input JSON", inputJSON).Msg("Input JSON")
		pretty, err := prettyJson(inputJSON)
		if err != nil {
			log.Err(err).Msg("Failed to pretty json")
			fmt.Printf("unable to convert to json %s\n", err)
			return err.Error()
		}
		return pretty
	})
	return jsonFunc
}
