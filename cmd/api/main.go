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
		Addr: env.GetString(config.ADDR, config.DEFAULT_PORT),
		DB: components.DBConfig{
			Addr:            env.GetString(config.DB_ADDR, config.DEFAULT_DB_ADDR),
			MaxOpenConns:    env.GetInt(config.DB_MAX_OPEN_CONNS, config.DEFAULT_MAX_OPEN_CONNS),
			MaxIdleConns:    env.GetInt(config.DB_MAX_IDLE_CONNS, config.DEFAULT_MAX_IDLE_CONNS),
			MaxIdleTimeMins: env.GetString(config.DB_CONN_MAX_IDLE_TIME, config.DEFAULT_DB_CONN_MAX_IDLE_TIME),
		},
	}
	dbConn, err := db.New(cfg.DB.Addr, cfg.DB.MaxOpenConns, cfg.DB.MaxIdleConns, cfg.DB.MaxIdleTimeMins)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer dbConn.Close()
	store := store.NewStorage(dbConn)
	app := &components.Application{
		Config: cfg,
		Store:  store,
	}
	mux := app.Mount()
	if err := app.Run(mux); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
