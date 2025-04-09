package config

// ключ и значение для порта HTTP-сервера
const (
	ADDR         = "ADDR"
	DEFAULT_PORT = ":8080"
)

// ключ и значение для адреса БД
const (
	DB_ADDR         = "DB_ADDR"
	DEFAULT_DB_ADDR = "postgres://postgres:545687@localhost/networkdb?sslmode=disable"
)

// ключ и значение для максимального числа открытых соединений
const (
	DB_MAX_OPEN_CONNS      = "DB_MAX_OPEN_CONNS"
	DEFAULT_MAX_OPEN_CONNS = 100
)

// ключ и значение для максимального числа простаивающих соединений
const (
	DB_MAX_IDLE_CONNS      = "DB_MAX_IDLE_CONNS"
	DEFAULT_MAX_IDLE_CONNS = 100
)

// ключ и значение для времени простоя соединения (строка, которую Env парсит в Duration)
const (
	DB_MAX_IDLE_TIME      = "DB_MAX_IDLE_TIME"
	DEFAULT_MAX_IDLE_TIME = "15m"
)

const (
	QUERY_STR string = `INSERT INTO posts (user_id, title, content) VALUES ($1, $2, $3, $4) returning id, created_at, updated_at`
)
