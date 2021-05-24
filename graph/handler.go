package graph

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Antuans-Tavern/ecommerce-backend/graph/generated"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func QueryHandler(db *gorm.DB) echo.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			DB: db,
		},
	}))

	// h.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
	// 	err := graphql.DefaultErrorPresenter(ctx, e)

	// 	fmt.Println("---------------------" + err.Error() + "-----------------")

	// 	return nil
	// })

	return func(c echo.Context) error {
		h.ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}
}
