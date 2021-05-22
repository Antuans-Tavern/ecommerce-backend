package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/Antuans-Tavern/ecommerce-backend/graph/generated"
	"github.com/Antuans-Tavern/ecommerce-backend/graph/types"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database/model"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/util"
)

func (r *queryResolver) Login(ctx context.Context, email string, password string) (*types.Login, error) {
	user := &model.User{}
	if err := r.DB.WithContext(ctx).Joins("Profile").First(user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	if !util.CompareHash(user.Password, password) {
		return nil, errors.New("123123123")
	}

	accessToken, err := user.CreateAccressToken()

	if err != nil {
		return nil, err
	}

	return &types.Login{
		Token: accessToken,
		User:  user,
	}, nil
}

func (r *queryResolver) Register(ctx context.Context, data types.Register) (*types.Login, error) {
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

	accessToken, err := user.CreateAccressToken()

	if err != nil {
		return nil, err
	}

	return &types.Login{
		Token: accessToken,
		User:  user,
	}, nil

}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
