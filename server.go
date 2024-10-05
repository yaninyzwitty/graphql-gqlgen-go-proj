package main

import (
	"log/slog"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yaninyzwitty/graphql-ggqlen-go-proj/configuration"
	"github.com/yaninyzwitty/graphql-ggqlen-go-proj/database"
	"github.com/yaninyzwitty/graphql-ggqlen-go-proj/graph"
)

func main() {
	cfg, err := configuration.LoadConfig()
	if err != nil {
		slog.Error("Failed to load configuration", "error", err)
		return
	}

	db, err := database.NewDatabaseConnection(cfg.DB_URL)
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		return
	}

	defer db.Close()
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	resolver := &graph.Resolver{
		DB: db,
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	slog.Info("listening on :" + cfg.PORT)
	err = http.ListenAndServe(":"+cfg.PORT, router)
	if err != nil {
		return
	}
}
