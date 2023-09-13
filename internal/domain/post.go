package domain

import (
	"context"
	"time"
)

// model
type PostID int64

type Post struct {
	ID        PostID    `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type Posts []*User

// repository interface
type PostRepository interface {
	AddPost(ctx context.Context, post *Post) error
	// FindPosts(ctx context.Context) (Posts, error)
}
