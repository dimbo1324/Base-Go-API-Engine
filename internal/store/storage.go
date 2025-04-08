package store

import (
	"database/sql"
)

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db},
		Users: &UsersStore{db},
	}
}
