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

	code := m.Run()
	os.Exit(code)
}

func TestProductsQuery(t *testing.T) {
	db, err := gorm.Open(postgres.Open(database.ConnectionString()), &gorm.Config{})
	assert.Nil(t, err)

	seeder.Call(db,
		seeder.CategorySeeder,
		seeder.ProductSeeder,
	)

	categoryID := int(1)

	products, err := ProductResolver(db, context.Background(), nil, 10, 1, &categoryID)
	assert.Nil(t, err)

	assert.Equal(t, 10, len(products))
	assert.Equal(t, uint(categoryID), products[0].CategoryID)
}
