package resolver

import (
	"context"
	"os"
	"testing"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database/seeder"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/util"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	viper.AddConfigPath("../..")
	util.PrepareDatabase()

	db, _ := gorm.Open(postgres.Open(database.ConnectionString()), &gorm.Config{})

	seeder.Call(db,
		seeder.CategorySeeder,
		seeder.ProductSeeder,
	)

	code := m.Run()
	os.Exit(code)
}

func TestProductsQuery(t *testing.T) {
	db, err := gorm.Open(postgres.Open(database.ConnectionString()), &gorm.Config{})
	assert.Nil(t, err)

	categoryID := int(1)

	collection, err := ProductsResolver(db, context.Background(), nil, 10, 1, &categoryID)
	assert.Nil(t, err)

	assert.Equal(t, 10, len(collection.Data))
	assert.Equal(t, uint(categoryID), collection.Data[0].CategoryID)
}

func TestProductQuery(t *testing.T) {
	db, err := gorm.Open(postgres.Open(database.ConnectionString()), &gorm.Config{})
	assert.Nil(t, err)

	productID := int(1)

	product, err := ProductResolver(db, context.Background(), productID)
	assert.Nil(t, err)

	assert.NotEmpty(t, product)
}
