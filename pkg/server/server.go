package server

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/graph"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/graphqlcontext"
	"github.com/fatih/color"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// SetUp configures and runs the webserver
func SetUp(db *gorm.DB) {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(graphqlcontext.Middleware)

	r.Get("/", playground.Handler("GraphiQL", "/query"))
	r.Post("/query", graph.QueryHandler(db))

	port := viper.GetString("port")
	color.Blue("Listening on port: %s", port)
	http.ListenAndServe(":"+port, r)
}
