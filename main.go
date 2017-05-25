package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dreae/esi-graphql/resolvers"
	rhttp "github.com/dreae/esi-graphql/resolvers/http"
	graphql "github.com/neelance/graphql-go"
)

var schema *graphql.Schema
var rootDir = "."

func buildSchema() (string, error) {
	schemaFiles, err := AssetDir("assets/schema")
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	for _, file := range schemaFiles {
		contents, err := Asset("assets/schema/" + file)
		if err != nil {
			return "", err
		}

		if _, err := buffer.Write(contents); err != nil {
			return "", err
		}
	}

	return string(buffer.Bytes()), nil
}

func init() {
	var err error
	schemaFile, err := buildSchema()
	if err != nil {
		panic(err)
	}

	schema, err = graphql.ParseSchema(schemaFile, &resolvers.Resolver{})
	if err != nil {
		panic(err)
	}

	page, err := Asset("assets/index.html")
	if err != nil {
		panic("Could not read index file from ./assets/index.html")
	}

	rhttp.InitHTTP()

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
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
