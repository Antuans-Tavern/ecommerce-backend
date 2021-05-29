package main

import (
	"github.com/Antuans-Tavern/ecommerce-backend/config"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database/model"
	_ "github.com/Antuans-Tavern/ecommerce-backend/pkg/lang"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/server"
	"github.com/spf13/viper"
)

func main() {
	config.Load()
	db, err := database.Connect()

	if err != nil {
		panic(err)
	}

	database.Migrate(db)
	model.SetPivots(db)

	e := server.SetUp(db)

	port := viper.GetString("port")
	e.Logger.Fatal(e.Start(":" + port))
}
