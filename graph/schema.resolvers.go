package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/Antuans-Tavern/ecommerce-backend/graph/generated"
	"github.com/Antuans-Tavern/ecommerce-backend/graph/types"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database/model"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/lang"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/resolver"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/util"
)

func (r *mutationResolver) Register(ctx context.Context, data types.Register) (*types.Login, error) {
	user := &model.User{
		Email:    data.Email,
		Password: util.Hash(data.Password),
		Profile: model.Profile{
			Name:     data.Name,
			Lastname: data.Lastname,
		},
	}

	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}

	accessToken, err := user.CreateAccressToken(r.DB)

	if err != nil {
		return nil, err
	}

	return &types.Login{
		Token: accessToken,
		User:  user,
	}, nil
}

func (r *queryResolver) Login(ctx context.Context, email string, password string) (*types.Login, error) {
	user := &model.User{}
	if err := r.DB.WithContext(ctx).Joins("Profile").First(user, "email = ?", email).Error; err != nil {
		return nil, errors.New(lang.Translator().Sprintf("login error"))
	}

	if !util.CompareHash(user.Password, password) {
		return nil, errors.New(lang.Translator().Sprintf("login error"))
	}

	accessToken, err := user.CreateAccressToken(r.DB)

	if err != nil {
		return nil, err
	}

	return &types.Login{
		Token: accessToken,
		User:  user,
	}, nil
}

func (r *queryResolver) Products(ctx context.Context, pagination int, page int) ([]*model.Product, error) {
	return resolver.ProductResolver(r.DB, ctx, pagination, page)
}

func (r *queryResolver) SearchProduct(ctx context.Context, search string, pagination int, page int) ([]*model.Product, error) {
	return resolver.ProductSearch(r.DB, ctx, search, pagination, page)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
