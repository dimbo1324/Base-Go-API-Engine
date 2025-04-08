package store

import (
	"context"
	"database/sql"

)

type PostStore struct {
	db *sql.DB
}
type Post struct {
	Id        int64    `json:"id"`
	UserId    int64    `json:"user_id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	Tags      []string `json:"tags"`
}
type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
}
type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
	}
	Users interface {
		Create(context.Context, *User) error
	}
}
type UsersStore struct {
	db *sql.DB
}
