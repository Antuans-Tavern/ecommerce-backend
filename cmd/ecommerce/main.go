package main

import (
	"github.com/Antuans-Tavern/ecommerce-backend/config"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database/model"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/server"
)

func main() {
	config.Load()

	db, err := database.Connect()

	if err != nil {
		panic(err)
	}

	database.Migrate(db)
	model.SetPivots(db)

	server.Serve(db)
}
