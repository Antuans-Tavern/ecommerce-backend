package graph

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/graph/generated"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func QueryHandler(db *gorm.DB) echo.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			DB: db,
		},
	}))

	return func(c echo.Context) error {
		h.ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}
}
