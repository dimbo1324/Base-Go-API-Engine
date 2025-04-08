package config

const (
	PORT          string = ":8080"
	ADDR          string = "ADDR"
	DB_ADDR       string = "DB_ADDR"
	DB_WAY        string = "postgres://admin:admin_password@localhost/go_api_db?sslmode=disable"
	OPEN          string = "DB_MAX_OPEN_CONNS"
	IDLE          string = "DB_MAX_IDLE_CONNS"
	IDLE_TIME     string = "DB_MAX_IDLE_TIME_MINS"
	OPEN_VAL      int    = 100
	IDLE_VAL      int    = 100
	IDLE_TIME_VAL string = "15m"
	QUERY_STR     string = `INSERT INTO posts (user_id, title, content) VALUES ($1, $2, $3, $4) returning id, created_at, updated_at`
)
