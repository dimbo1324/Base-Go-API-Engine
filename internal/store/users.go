package store

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

func (s *UsersStore) Create(ctx context.Context, user *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	query := `
        INSERT INTO users (username, password, email)
        VALUES ($1, $2, $3)
        RETURNING id, created_at
    `
	err = s.db.QueryRowContext(ctx, query, user.Username, hashedPassword, user.Email).
		Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
