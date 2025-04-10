package main

import (
	"log"

	"github.com/dimbo1324/Base-Go-API-Engine/cmd/api/components"
	"github.com/dimbo1324/Base-Go-API-Engine/internal/config"
	"github.com/dimbo1324/Base-Go-API-Engine/internal/db"
	"github.com/dimbo1324/Base-Go-API-Engine/internal/env"
	"github.com/dimbo1324/Base-Go-API-Engine/internal/store"
)

func main() {
	cfg := components.Config{
		Addr: env.GetString(config.ADDR, config.PORT),
		DB: components.DBConfig{
			Addr:            env.GetString(config.DB_ADDR, config.DB_ADDR_VAL),
			MaxOpenConns:    env.GetInt(config.OPEN, config.OPEN_VAL),
			MaxIdleConns:    env.GetInt(config.IDLE, config.IDLE_VAL),
			MaxIdleTimeMins: env.GetString(config.IDLE_TIME, config.IDLE_TIME_VAL),
		},
	}
	db, err := db.New(cfg.DB.Addr, cfg.DB.MaxOpenConns, cfg.DB.MaxIdleConns, cfg.DB.MaxIdleTimeMins)
	if err != nil {
		log.Panic(err)
	}
	store := store.NewStorage(db)
	app := &components.Application{
		Config: cfg,
		Store:  store,
	}
	mux := app.Mount()
	log.Fatal(app.Run(mux))
}
