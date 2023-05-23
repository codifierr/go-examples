package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/parser"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Info().Msg("Initializing app...")

	query, err := getQuery("../queries/country.graphql")
	if err != nil {
		log.Error().AnErr("Error", err).Msg("Error in getting query")
		return
	}

	varTypeMap, err := parseQuery(query)
	if err != nil {
		log.Error().AnErr("Error", err).Msg("Error in parsing query")
	}

	fmt.Println(varTypeMap)

}

func parseQuery(query string) (map[string]string, error) {
	varTypeMap := make(map[string]string)

	doc, err := parser.ParseQuery(&ast.Source{Input: query})
	if err != nil {
		return varTypeMap, err
	}
	for _, v := range doc.Operations {

		for _, variable := range v.VariableDefinitions {
			varTypeMap[variable.Variable] = variable.Type.NamedType
		}
	}
	return varTypeMap, nil
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
