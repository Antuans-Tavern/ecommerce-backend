package resolver

import (
	"context"
	"testing"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCategoriesQuery(t *testing.T) {
	db, err := gorm.Open(postgres.Open(database.ConnectionString()), &gorm.Config{})
	assert.Nil(t, err)

	collection, err := CategoriesResolver(db, context.Background(), nil, 10, 1, 10, 1)
	assert.Nil(t, err)

	assert.Equal(t, 10, len(collection.Data))
	assert.Equal(t, 10, len(collection.Data[0].Products))
}

func TestCategoryQuery(t *testing.T) {
	db, err := gorm.Open(postgres.Open(database.ConnectionString()), &gorm.Config{})
	assert.Nil(t, err)

	categoryID := int(1)

	category, err := CategoryResolver(db, context.Background(), categoryID, 10, 1)
	assert.Nil(t, err)

	assert.NotEmpty(t, category)
}
