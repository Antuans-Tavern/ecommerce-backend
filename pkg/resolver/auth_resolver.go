package resolver

import (
	"context"
	"errors"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database/model"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/graph/types"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/graphqlcontext"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/lang"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/util"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type input struct {
	Email string `validate:"email"`
}

// Register resolver
func Register(db *gorm.DB, validate *validator.Validate, ctx context.Context, data types.Register) (*types.Login, error) {
	user := model.User{
		Email:    data.Email,
		Password: util.Hash(data.Password),
		Profile: model.Profile{
			Name:     data.Name,
			Lastname: data.Lastname,
		},
	}

	var err error
	if err = validate.Struct(user); err != nil {
		return nil, util.ValidationError(ctx, err.(validator.ValidationErrors))
	}

	var accessToken string

	if accessToken, err = user.Register(ctx, db, data.Authenticate); err != nil {
		return nil, err
	}

	return &types.Login{
		Token: accessToken,
		User:  &user,
	}, nil
}

func Login(db *gorm.DB, validate *validator.Validate, ctx context.Context, email string, password string) (*types.Login, error) {

	if errs := validate.Struct(input{Email: email}); errs != nil {
		return nil, util.ValidationError(ctx, errs.(validator.ValidationErrors))
	}

	var err error

	user := &model.User{}
	if err = db.First(user, "email = ?", email).Error; err == gorm.ErrRecordNotFound {
		return nil, errors.New(lang.Trans(graphqlcontext.GetContext(ctx).Locale, "login error"))
	}

	if !util.CompareHash(user.Password, password) {
		return nil, errors.New(lang.Trans(graphqlcontext.GetContext(ctx).Locale, "login error"))
	}

	var accessToken string
	if accessToken, err = user.CreateAccressToken(db, graphqlcontext.GetContext(ctx).UserAgent); err != nil {
		return nil, err
	}

	return &types.Login{
		Token: accessToken,
		User:  user,
	}, nil
}
