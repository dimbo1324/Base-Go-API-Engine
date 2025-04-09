package main

import (
	"log"

	"github.com/example/universal-api-engine/cmd/api/components"
	"github.com/example/universal-api-engine/internal/config"
	"github.com/example/universal-api-engine/internal/db"
	"github.com/example/universal-api-engine/internal/env"
	"github.com/example/universal-api-engine/internal/store"
)

// Главная точка входа в приложение
func main() {
	cfg := components.Config{
		Addr: env.GetString(config.ADDR, config.DEFAULT_PORT),
		DB: components.DBConfig{
			Addr:            env.GetString(config.DB_ADDR, config.DEFAULT_DB_ADDR),
			MaxOpenConns:    env.GetInt(config.DB_MAX_OPEN_CONNS, config.DEFAULT_MAX_OPEN_CONNS),
			MaxIdleConns:    env.GetInt(config.DB_MAX_IDLE_CONNS, config.DEFAULT_MAX_IDLE_CONNS),
			MaxIdleTimeMins: env.GetString(config.DB_MAX_IDLE_TIME, config.DEFAULT_MAX_IDLE_TIME),
		},
	}
	// Подключение к базе данных с обработкой ошибок
	dbConn, err := db.New(cfg.DB.Addr, cfg.DB.MaxOpenConns, cfg.DB.MaxIdleConns, cfg.DB.MaxIdleTimeMins)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer dbConn.Close()
	// Инициализация хранилища данных
	store := store.NewStorage(dbConn)
	// Создание приложения
	app := &components.Application{
		Config: cfg,
		Store:  store,
	}
	// Настройка маршрутов и запуск сервера
	mux := app.Mount()
	if err := app.Run(mux); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
