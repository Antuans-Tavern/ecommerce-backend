package graph

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/lang"
	"github.com/spf13/viper"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func recover(ctx context.Context, err interface{}) error {

	if viper.GetBool("debug") {
		graphql.AddErrorf(ctx, "%v", err)
	}

	return errors.New(lang.Translator().Sprintf("internal error"))
}

func errorPresenter(ctx context.Context, e error) *gqlerror.Error {

	// if err, ok := e.(validator.ValidationErrors); ok {
	// }

	err := graphql.DefaultErrorPresenter(ctx, e)

	return err
}
