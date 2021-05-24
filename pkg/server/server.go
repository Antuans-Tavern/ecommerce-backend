package server

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Antuans-Tavern/ecommerce-backend/graph"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func playgroundHandler() echo.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c echo.Context) error {
		h.ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}
}

func SetUp(db *gorm.DB) *echo.Echo {
	e := echo.New()

	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(60)))

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogLevel: log.INFO,
	}))

	e.POST("/query", graph.QueryHandler(db))
	e.GET("/", playgroundHandler())

	return e
}
