package main

import (
	"fmt"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"io/ioutil"
	"log"
	"net/http"
	models "github.com/bezaeel/fynd-twitter/models"
)



var opts =  []graphql.SchemaOpt{graphql.UseFieldResolvers()}
var posts = []models.Post{
	{
		Id: "0",
		Message: "Ask Talabi..",
	},
}


type RootResolver struct {

}



func (r *RootResolver) Ping() (string, error) {
	return "Ask Talabi..", nil
}

func (r *RootResolver) Posts() ([]models.Post, error){
	return posts, nil
}


func parseSchema(path string, resolver interface{}) *graphql.Schema {
	bstr, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	schemaString := string(bstr)
	parsedSchema, err := graphql.ParseSchema(schemaString, resolver, opts...,)

	if err != nil {
		panic(err)
	}
	return parsedSchema
}


func main(){
	http.Handle("/graphql", &relay.Handler{
		Schema: parseSchema("./schema.graphql", &RootResolver{}),
	})

	fmt.Println("serving on 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}