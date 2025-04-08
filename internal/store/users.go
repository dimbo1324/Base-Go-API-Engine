package store

import (
	"context"

)

func (s *UsersStore) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (username, password, email)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	return s.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).
		Scan(&user.ID, &user.CreatedAt)
}
