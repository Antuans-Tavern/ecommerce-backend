package util

import (
	"database/sql"
	"fmt"

	"github.com/Antuans-Tavern/ecommerce-backend/config"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func PrepareDatabase() {
	config.Load()

	createDb()

	viper.Set("db_name", viper.GetString("db_name_test"))
	orm, _ := database.Connect()

	database.Drop(orm)
	database.Migrate(orm)

}

func createDb() {
	if db, err := sql.Open("postgres", database.TestingConnectionString()); err != nil || db.Ping() != nil {
		if err != nil {
			panic(err)
		}

		db.Close()

		db, _ = sql.Open("postgres", database.ConnectionString())
		defer db.Close()
		db.Exec(fmt.Sprintf("CREATE DATABASE %s;", viper.GetString("db_name_test")))
	}
}
