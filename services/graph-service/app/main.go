package main

import (
	"log"
	"net/http"
	"os"

	"github.com/0x726f6f6b6965/mono-service/services/graph-service/graph"
	"github.com/0x726f6f6b6965/mono-service/services/graph-service/internal/middleware"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	var handler *chi.Mux = NewGraphQLHandler()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

// NewGraphQLHandler returns handler for GraphQL application
func NewGraphQLHandler() *chi.Mux {
	// create a new router
	var router *chi.Mux = chi.NewRouter()

	// use the middleware component
	router.Use(middleware.AuthMiddleware())

	// create a GraphQL server
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// assign some handlers for the GraphQL server
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	// return the handler
	return router
}
