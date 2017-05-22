package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"context"

	"github.com/dreae/esi-graphql/resolvers"
	"github.com/neelance/graphql-go"
)

var schema *graphql.Schema

func init() {
	var err error
	schemaFile, err := Asset("assets/schema.gql")
	if err != nil {
		panic(err)
	}

	schema, err = graphql.ParseSchema(string(schemaFile), &resolvers.Resolver{})
	if err != nil {
		panic(err)
	}
}

func main() {
	page, err := Asset("assets/index.html")
	if err != nil {
		panic("Could not read index file from ./assets/index.html")
	}

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	}))

	http.Handle("/query", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var params struct {
			Query         string                 `json:"query"`
			OperationName string                 `json:"operationName"`
			Variables     map[string]interface{} `json:"variables"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), resolvers.ContextKey("auth"), r.Header.Get("Authorization"))

		response := schema.Exec(ctx, params.Query, params.OperationName, params.Variables)

		responseJSON, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(responseJSON)
	}))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
