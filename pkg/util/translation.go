package util

import (
	"context"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/graphqlcontext"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/lang"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func ValidationError(ctx context.Context, errs validator.ValidationErrors) error {
	extensions := make(map[string]interface{}, len(errs))

	locale := ctx.Value("context").(graphqlcontext.Context).Locale

	for key, value := range errs.Translate(lang.Translator(locale)) {
		extensions[key] = value
	}

	return &gqlerror.Error{
		Message:    "Validation error",
		Extensions: extensions,
	}

}

func GetTranslator() ut.Translator {
	return viper.Get("trans").(ut.Translator)
}
