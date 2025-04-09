package store

import (
	"context"

	"github.com/dimbo1324/Base-Go-API-Engine/internal/config"
	"github.com/lib/pq"
)

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := config.QUERY_STR
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
