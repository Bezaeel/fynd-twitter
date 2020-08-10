package main

import (
	"encoding/json"
	"fmt"
	"github.com/bezaeel/fynd-twitter/resolvers"
	"github.com/graphql-go/graphql"
	"net/http"
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			/* Get (read) single post by id
			   http://localhost:8000/post?query={post(id:1){message,comment}}
			*/
			"post": resolvers.GetPost,
			/* Get (read) post list
			   http://localhost:8000/post?query={list{id,message,comment}}
			*/
			"posts": resolvers.ListPosts,
			"user":  resolvers.GetUser,
			"users": resolvers.ListUsers,
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
		//Mutation: mutationType,
	},
)

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
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Server is running on port 8000")
	http.ListenAndServe(":8000", nil)
}
