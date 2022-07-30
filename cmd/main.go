package main

import (
	"context"
	"log"
	"net/http"
	"testrealworld"

	"testrealworld/auth"
	"testrealworld/ent"
	"testrealworld/ent/migrate"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"

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

	router := chi.NewRouter()

	router.Use(auth.Middleware(client))

	// Configure the server and start listening on :8081.
	srv := handler.NewDefaultServer(testrealworld.NewSchema(client))
	router.Handle("/",
		playground.Handler("Todo", "/query"),
	)
	router.Handle("/query", srv)
	log.Println("listening on :8081")
	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Fatal("http server terminated", err)
	}
}
