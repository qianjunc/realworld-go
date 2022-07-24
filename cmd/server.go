package main

import (
	"context"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/qianjunc/realworld/ent"
	"github.com/qianjunc/realworld/graph"
	"github.com/qianjunc/realworld/graph/generated"
)

const defaultPort = "8080"

func main() {
	client, err := ent.Open("mysql", "root:dbpass@tcp(localhost:3306)/hackernews?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
