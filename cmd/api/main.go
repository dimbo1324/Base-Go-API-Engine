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
		Addr: env.GetString(config.KeyName, config.Port),
		DB: components.DBConfig{
			Addr:            env.GetString("DB_ADDR", "postgres://admin:admin_password@localhost/go_api_db?sslmode=disable"),
			MaxOpenConns:    env.GetInt("DB_MAX_OPEN_CONNS", 100),
			MaxIdleConns:    env.GetInt("DB_MAX_IDLE_CONNS", 100),
			MaxIdleTimeMins: env.GetString("DB_MAX_IDLE_TIME_MINS", "15"),
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
