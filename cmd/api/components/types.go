package components

import "github.com/dimbo1324/Base-Go-API-Engine/internal/store"

type Config struct {
	Addr string
}
type Application struct {
	Config Config
	Store  store.Storage
}
