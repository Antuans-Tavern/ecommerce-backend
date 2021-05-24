package database

import (
	"fmt"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database/model"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestingConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		viper.GetString("db_host"),
		viper.GetString("db_port"),
		viper.GetString("db_user"),
		viper.GetString("db_password"),
		viper.GetString("db_name_test"),
	)
}

func ConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		viper.GetString("db_host"),
		viper.GetString("db_port"),
		viper.GetString("db_user"),
		viper.GetString("db_password"),
		viper.GetString("db_name"),
	)
}

func Connect() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(ConnectionString()), &gorm.Config{
		PrepareStmt: true,
	})
}

func Models() []interface{} {
	return []interface{}{
		&model.User{},
		&model.Profile{},
		&model.Category{},
		&model.Product{},
		&model.ShoppingCart{},
		&model.ShoppingCartProduct{},
		&model.Configuration{},
		&model.Invoice{},
		&model.InvoiceItem{},
		&model.Payment{},
		&model.Image{},
		&model.Configuration{},
		&model.Coupon{},
		&model.AccessToken{},
	}
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		Models()...,
	)
}

func Drop(db *gorm.DB) {
	db.Migrator().DropTable(
		Models()...,
	)
}
