package graph

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/graphqlcontext"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/lang"
	"github.com/spf13/viper"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func recover(ctx context.Context, err interface{}) error {

	if viper.GetBool("debug") {
		graphql.AddErrorf(ctx, "%v", err)
	}

	locale := ctx.Value("context").(graphqlcontext.Context).Locale

	return errors.New(lang.Trans(locale, "internal error"))
}

func errorPresenter(ctx context.Context, e error) *gqlerror.Error {
	err := graphql.DefaultErrorPresenter(ctx, e)
	return err
}
