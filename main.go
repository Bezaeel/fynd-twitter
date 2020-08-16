package main

import (
	"encoding/json"
	"fmt"
	"github.com/bezaeel/fynd-twitter/resolvers"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"post":  resolvers.GetPost,
			"posts": resolvers.ListPosts,
			"user":  resolvers.GetUser,
			"users": resolvers.ListUsers,
		},
	})

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"AddFollower": resolvers.AddFollower,
		"AddPost":     resolvers.AddPost,
	},
})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)
var h = handler.New(&handler.Config{
	Schema: &schema,
	Pretty: true,
})

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func main() {

	http.Handle("/fynd", h)
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		log.Println("request: ", r.Body)
		result := executeQuery(r.URL.Query().Get("mutation"), schema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Server is running on port 8000")
	http.ListenAndServe(":8000", nil)
}
