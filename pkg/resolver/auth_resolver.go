package resolver

import (
	"context"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database/model"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/graph/types"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/util"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func Register(db *gorm.DB, validator *validator.Validate, ctx context.Context, data types.Register) (*types.Login, error) {
	user := model.User{
		Email:    data.Email,
		Password: util.Hash(data.Password),
		Profile: model.Profile{
			Name:     data.Name,
			Lastname: data.Lastname,
		},
	}

	var err error
	if err = validator.Struct(user); err != nil {
		return nil, err
	}

	var accessToken string

	if accessToken, err = user.Register(ctx, db); err != nil {
		return nil, err
	}

	return &types.Login{
		Token: accessToken,
		User:  &user,
	}, nil
}
