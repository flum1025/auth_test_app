package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/flum1025/test_auth_app/api"
	middleware "github.com/flum1025/test_auth_app/api/server/middleware"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	origin := os.Getenv("ORIGIN")

	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{origin},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.Use(middleware.GenerateMiddleware())
	router.Use(middleware.GqlLogMiddleware)
	router.Handle("/query", handler.GraphQL(api.NewExecutableSchema(api.Config{Resolvers: &api.Resolver{}})))

	log.Printf("connect to http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
