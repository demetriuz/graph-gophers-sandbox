package main

import (
	"io/ioutil"
	"log"
	"net/http"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	queryResolvers "github.com/demetriuz/graph-gophers-sandbox/query_resolvers"
)


func main() {
	s, err := getSchema("./schema.graphql")
	if err != nil {
		panic(err)
	}

	schema := graphql.MustParseSchema(s, &queryResolvers.Query{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}


func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}