package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Load() {

	viper.SetDefault("app_secret", "secret")

	viper.SetDefault("lang", "es")
	viper.SetDefault("timezone", "UTC")

	viper.SetDefault("port", 80)
	viper.SetDefault("debug", false)

	viper.SetDefault("db_name_test", "test_db")
	viper.SetDefault("db_port", 5432)

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Config file error: %s \n", err)
	}

}
