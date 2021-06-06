package graph

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/graph/generated"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/rule"
	"gorm.io/gorm"
)

func QueryHandler(db *gorm.DB) http.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			DB:        db,
			Validator: rule.SetUpValidator(db),
		},
	}))

	h.SetRecoverFunc(recover)
	h.SetErrorPresenter(errorPresenter)

	return h.ServeHTTP
}
