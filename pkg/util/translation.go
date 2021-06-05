package util

import (
	"context"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/graphqlcontext"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/lang"
	"github.com/go-playground/validator/v10"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func ValidationError(ctx context.Context, errs validator.ValidationErrors) error {
	extensions := make(map[string]interface{}, len(errs))

	locale := graphqlcontext.GetContext(ctx).Locale

	for key, value := range errs.Translate(lang.Translator(locale)) {
		extensions[key] = value
	}

	return &gqlerror.Error{
		Message:    "Validation error",
		Extensions: extensions,
	}

}
