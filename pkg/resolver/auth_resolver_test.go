package resolver

import (
	"context"
	"testing"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/graph/types"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/graphqlcontext"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/rule"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestRegisterResolver(t *testing.T) {
	db, err := gorm.Open(postgres.Open(database.ConnectionString()), &gorm.Config{})
	assert.Nil(t, err)

	validate := rule.SetUpValidator(db)

	data := types.Register{
		Email:    gofakeit.Email(),
		Password: "secret123",
		Name:     gofakeit.Name(),
		Lastname: gofakeit.LastName(),
	}

	login, err := Register(db, validate, context.Background(), data)
	assert.Nil(t, err)
	assert.NotNil(t, login)
	assert.NotNil(t, login.User)
	assert.NotNil(t, login.User.Profile)
	assert.Empty(t, login.Token)
}

func TestRegisterResolverWithoutAuthenticate(t *testing.T) {
	db, err := gorm.Open(postgres.Open(database.ConnectionString()), &gorm.Config{})
	assert.Nil(t, err)

	validate := rule.SetUpValidator(db)

	data := types.Register{
		Email:        gofakeit.Email(),
		Password:     "secret123",
		Name:         gofakeit.Name(),
		Lastname:     gofakeit.LastName(),
		Authenticate: false,
	}

	login, err := Register(db, validate, context.Background(), data)
	assert.Nil(t, err)
	assert.NotNil(t, login)
	assert.NotNil(t, login.User)
	assert.NotNil(t, login.User.Profile)
	assert.Empty(t, login.Token)
}

func TestLogin(t *testing.T) {
	db, err := gorm.Open(postgres.Open(database.ConnectionString()), &gorm.Config{})
	assert.Nil(t, err)

	validate := rule.SetUpValidator(db)

	login, err := Login(db, validate, context.WithValue(context.Background(), "context", graphqlcontext.Context{}), "admin@testing.com", "secret")
	assert.Nil(t, err)
	assert.NotNil(t, login)
	assert.NotNil(t, login.User)
	assert.NotNil(t, login.User.Profile)
	assert.NotEmpty(t, login.Token)
}
