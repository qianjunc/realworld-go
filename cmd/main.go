package main

import (
	"context"
	"log"
	"net/http"
	"testrealworld"

	"testrealworld/ent"
	"testrealworld/ent/migrate"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Create ent.Client and run the schema migration.

	client, err := ent.Open("mysql", "root:dbpass@tcp(localhost:3306)/hackernews?parseTime=True")
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	// Configure the server and start listening on :8081.
	srv := handler.NewDefaultServer(testrealworld.NewSchema(client))
	http.Handle("/",
		playground.Handler("Todo", "/query"),
	)
	http.Handle("/query", srv)
	log.Println("listening on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}
