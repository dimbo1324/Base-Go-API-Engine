package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type Post struct {
	Id        int64    `json: "id"`
	UserId    int64    `json: "user_id"`
	Content   string   `json: "content"`
	Title     string   `json: "title"`
	CreatedAt string   `json: "created_at"`
	UpdatedAt string   `json: "updated_at"`
	Tags      []string `json: "tags"`
}
type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `INSERT INTO posts (user_id, title, content) VALUES ($1, $2, $3, $4) returning id, created_at, updated_at`
	err := s.db.QueryRowContext(
		ctx,
		query,
		post.UserId,
		post.Title,
		post.Content,
		pq.Array(post.Tags),
	).Scan(
		&post.Id,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
