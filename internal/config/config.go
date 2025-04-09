package config

const (
	ADDR         = "ADDR"
	DEFAULT_PORT = ":8080"
)
const (
	DB_ADDR         = "DB_WAY"
	DEFAULT_DB_ADDR = "postgres://postgres:545687@localhost/appdb?sslmode=disable"
)
const (
	DB_MAX_OPEN_CONNS      = "DB_MAX_OPEN_CONNS"
	DEFAULT_MAX_OPEN_CONNS = 100
)
const (
	DB_MAX_IDLE_CONNS      = "DB_MAX_IDLE_CONNS"
	DEFAULT_MAX_IDLE_CONNS = 100
)
const (
	DB_MAX_IDLE_TIME      = "DB_MAX_IDLE_TIME"
	DEFAULT_MAX_IDLE_TIME = "15m"
)
const (
	QUERY_STR string = `INSERT INTO posts (user_id, title, content, tags) VALUES ($1, $2, $3, $4) returning id, created_at, updated_at`
)
