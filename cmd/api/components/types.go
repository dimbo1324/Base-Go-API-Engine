package components

import "github.com/dimbo1324/Base-Go-API-Engine/internal/store"

type DBConfig struct {
	Addr            string
	MaxOpenConns    int
	MaxIdleConns    int
	MaxIdleTimeMins string
}
type Config struct {
	Addr string
	DB   DBConfig
}
type Application struct {
	Config Config
	Store  store.Storage
}
