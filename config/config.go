package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Load() {

	viper.SetDefault("port", 80)
	viper.SetDefault("timezone", "UTC")
	viper.SetDefault("debug", false)

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Config file error: %s \n", err)
	}

}
