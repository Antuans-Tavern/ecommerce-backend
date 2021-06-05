package graph

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/graph/generated"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/rule"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func QueryHandler(db *gorm.DB) http.HandlerFunc {
	validate := validator.New()
	en := en.New()
	es := es.New()
	uni := ut.New(es, en)

	trans, _ := uni.GetTranslator("es")

	validate.RegisterTranslation("db_unique", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})

	validate.RegisterValidation("db_unique", rule.DBUnique(db))

	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			DB:        db,
			Validator: validate,
		},
	}))

	h.SetRecoverFunc(recover)
	h.SetErrorPresenter(errorPresenter)

	return h.ServeHTTP
}
