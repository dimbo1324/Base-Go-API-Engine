package main

import (
	"log"

	"github.com/dimbo1324/Base-Go-API-Engine/cmd/api/components"
	"github.com/dimbo1324/Base-Go-API-Engine/internal/config"
	"github.com/dimbo1324/Base-Go-API-Engine/internal/env"
	"github.com/dimbo1324/Base-Go-API-Engine/internal/store"
)

func main() {
	cfg := components.Config{
		Addr: env.GetString(config.KeyName, config.Port),
	}
	store := store.NewStorage(nil)
	app := &components.Application{
		Config: cfg,
		Store:  store,
	}
	mux := app.Mount()
	log.Fatal(app.Run(mux))
}
