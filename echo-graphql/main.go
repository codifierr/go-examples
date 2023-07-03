package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var users []User = []User{
	{ID: "1", Name: "John"},
	{ID: "2", Name: "Jane"},
	{ID: "3", Name: "Bob"},
}

func main() {
	// create a new Echo instance
	e := echo.New()

	// add middleware for logging
	e.Use(middleware.Logger())

	// create a new GraphQL schema
	fields := graphql.Fields{
		"user": &graphql.Field{
			Type: graphql.NewObject(graphql.ObjectConfig{
				Name: "User",
				Fields: graphql.Fields{
					"id": &graphql.Field{
						Type: graphql.String,
					},
					"name": &graphql.Field{
						Type: graphql.String,
					},
				},
			}),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(string)
				if ok {
					for _, user := range users {
						if user.ID == id {
							return user, nil
						}
					}
				}
				return nil, nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create schema: %v", err)
	}

	// define a handler function for the GraphQL endpoint
	handler := func(c echo.Context) error {
		var params struct {
			Query string `json:"query"`
		}
		err := json.NewDecoder(c.Request().Body).Decode(&params)
		if err != nil {
			return err
		}
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: params.Query,
		})
		if len(result.Errors) > 0 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("failed to execute graphql operation, errors: %v", result.Errors))
		}
		return c.JSON(http.StatusOK, result)
	}

	// add a route for the GraphQL endpoint
	e.POST("/graphql", handler)

	// start the server
	e.Logger.Fatal(e.Start(":8080"))
}
