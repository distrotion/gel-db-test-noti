package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"mindfit.noti/graph"
	"mindfit.noti/graph/generated"
	"mindfit.noti/mongo/maindb"
)

const defaultPort = "9111"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	maindb.Getcol()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
