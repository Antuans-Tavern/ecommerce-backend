package main

import (
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
)

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if os.IsNotExist(err) {
		cfg, err = config.LoadDefaultConfig()
	}

	if err != nil {
		panic(err)
	}

	if err = api.Generate(cfg); err != nil {
		panic(err)
	}
}
