package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"context"

	"github.com/dreae/esi-graphql/resolvers"
	"github.com/neelance/graphql-go"
)

var schema *graphql.Schema

func init() {
	var err error
	schemaFile, err := ioutil.ReadFile("./assets/schema.gql")
	if err != nil {
		panic(err)
	}

	schema, err = graphql.ParseSchema(string(schemaFile), &resolvers.Resolver{})
	if err != nil {
		panic(err)
	}
}

func main() {
	page, err := ioutil.ReadFile("./assets/index.html")
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

		ctx := context.WithValue(r.Context(), "auth", r.Header.Get("Authorization"))

		response := schema.Exec(ctx, params.Query, params.OperationName, params.Variables)

		responseJSON, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(responseJSON)
	}))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
