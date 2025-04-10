package config

const (
	PORT          string = ":8080"
	ADDR          string = "ADDR"
	DB_ADDR       string = "DB_ADDR"
	DB_ADDR_VAL   string = "postgres://postgres:password@localhost/appdb?sslmode=disable"
	OPEN          string = "DB_MAX_OPEN_CONNS"
	IDLE          string = "DB_MAX_IDLE_CONNS"
	IDLE_TIME     string = "DB_MAX_IDLE_TIME_MINS"
	OPEN_VAL      int    = 100
	IDLE_VAL      int    = 100
	IDLE_TIME_VAL string = "15m"
	QUERY_STR     string = `  INSERT INTO posts (user_id, title, content, tags) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
`
)
