package resolver

import (
	"context"
	"errors"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database/model"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/graph/types"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/lang"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/util"
	"gorm.io/gorm"
)

func Register(db *gorm.DB, ctx context.Context, data types.Register) (*types.Login, error) {
	user := &model.User{
		Email:    data.Email,
		Password: util.Hash(data.Password),
		Profile: model.Profile{
			Name:     data.Name,
			Lastname: data.Lastname,
		},
	}

	if err := db.Create(user).Error; err != nil {
		return nil, err
	}

	accessToken, err := user.CreateAccressToken(db)

	if err != nil {
		return nil, err
	}

	return &types.Login{
		Token: accessToken,
		User:  user,
	}, nil
}

func Login(db *gorm.DB, ctx context.Context, email string, password string) (*types.Login, error) {
	user := &model.User{}
	if err := db.WithContext(ctx).Joins("Profile").First(user, "email = ?", email).Error; err != nil {
		return nil, errors.New(lang.Translator().Sprintf("login error"))
	}

	if !util.CompareHash(user.Password, password) {
		return nil, errors.New(lang.Translator().Sprintf("login error"))
	}

	accessToken, err := user.CreateAccressToken(db)

	if err != nil {
		return nil, err
	}

	return &types.Login{
		Token: accessToken,
		User:  user,
	}, nil
}
