package main

import (
	"log"

	"github.com/dimbo1324/Base-Go-API-Engine/cmd/api/components"
	"github.com/dimbo1324/Base-Go-API-Engine/internal/config"
	"github.com/dimbo1324/Base-Go-API-Engine/internal/env"
)

func main() {
	conf := components.Config{
		Addr: env.GetString(config.KeyName, config.Port),
	}

	app := &components.Application{
		Config: conf,
	}

	mux := app.Mount()
	log.Fatal(app.Run(mux))
}
