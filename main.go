package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

type Dogma struct {
	AttributeID  int    `json:"attribute_id"`
	DefaultValue int    `json:"default_value"`
	Description  string `json:"description"`
	DisplayName  string `json:"display_name"`
	HighIsGood   bool   `json:"high_is_good"`
	IconID       int    `json:"icon_id"`
	Name         string `json:"name"`
	Published    bool   `json:"published"`
	Stackable    bool   `json:"stackable"`
	UnitID       int    `json:"unit_id"`
}

func initGraphql() (graphql.Schema, error) {
	dogmaType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Dogma",
			// If only there were macros that could do this *grumble grumble*
			Fields: graphql.Fields{
				"AttributeID": &graphql.Field{
					Type: graphql.Int,
				},
				"DefaultValue": &graphql.Field{
					Type: graphql.Int,
				},
				"Description": &graphql.Field{
					Type: graphql.String,
				},
				"DisplayName": &graphql.Field{
					Type: graphql.String,
				},
				"HighIsGood": &graphql.Field{
					Type: graphql.Boolean,
				},
				"IconID": &graphql.Field{
					Type: graphql.Int,
				},
				"Name": &graphql.Field{
					Type: graphql.String,
				},
				"Published": &graphql.Field{
					Type: graphql.Boolean,
				},
				"Stackable": &graphql.Field{
					Type: graphql.Boolean,
				},
				"UnitID": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})

	queryType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"dogma": &graphql.Field{
					Type: dogmaType,
					Args: graphql.FieldConfigArgument{
						"AttributeID": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						id := p.Args["AttributeID"]
						resp, _ := http.Get(fmt.Sprintf("https://esi.tech.ccp.is/latest/dogma/attributes/%d/?datasource=tranquility", id))
						var dogma Dogma

						json.NewDecoder(resp.Body).Decode(&dogma)

						return dogma, nil
					},
				},
			},
		})

	return graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}

func main() {
	schema, _ := initGraphql()
	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: r.URL.Query()["query"][0],
		})
		json.NewEncoder(w).Encode(result)
	})

	http.ListenAndServe(":8080", http.DefaultServeMux)
}
