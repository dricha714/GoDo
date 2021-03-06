package main

import (
	_ "embed"
	"log"
	"net/http"

	"github.com/CodingProjects/Go/GoDo/common"
	"github.com/CodingProjects/Go/GoDo/controller"
	"github.com/friendsofgo/graphiql"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

//go:embed schema.gql
var schemaGQL []byte

func main() {

	log.Println("Creating TODO Directory...")
	err := common.CreateTodoDirectory()
	if err != nil {
		log.Printf("error occured while creating todo directory: ", err)
	}
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(string(schemaGQL), &controller.Controller{}, opts...)
	http.Handle("/query", &relay.Handler{Schema: schema})
	// init model
	// graphiql
	// First argument must be same as graphql handler path
	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/controller")
	if err != nil {
		panic(err)
	}
	http.Handle("/", graphiqlHandler)
	// Run
	log.Println("Server ready at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
