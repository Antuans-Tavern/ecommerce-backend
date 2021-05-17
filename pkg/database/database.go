package database

import (
	"fmt"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database/model"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		viper.GetString("db_host"),
		viper.GetString("db_port"),
		viper.GetString("db_user"),
		viper.GetString("db_password"),
		viper.GetString("db_name"),
	)
}

func Connect() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(connectionString()), &gorm.Config{
		PrepareStmt: true,
	})
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Profile{},
		&model.Category{},
		&model.Product{},
		&model.ShoppingCart{},
	)
}

func Drop(db *gorm.DB) {
	db.Migrator().DropTable(
		&model.User{},
		&model.Profile{},
		&model.Category{},
		&model.Product{},
		&model.ShoppingCart{},
	)
}
