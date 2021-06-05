package resolver

import (
	"os"
	"testing"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database/seeder"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/util"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	viper.AddConfigPath("../..")
	util.PrepareDatabase()

	db, _ := gorm.Open(postgres.Open(database.ConnectionString()), &gorm.Config{})

	seeder.Run(db)

	code := m.Run()
	os.Exit(code)
}
