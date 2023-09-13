package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/takumi2786/zero-backend/internal/domain"
)

type PostRepository struct {
	db *sqlx.DB
}

func (r *PostRepository) AddPost(ctx context.Context, post *domain.Post) error {
	sql := "INSERT INTO `posts` (" +
		"`title`, `content`, `updated_at`, `created_at` " +
		") " +
		"VALUES (" +
		"?, ?, NOW(), NOW()" +
		")"
	_, err := r.db.ExecContext(ctx, sql, post.Title, post.Content)
	if err != nil {
		return err
	}
	return nil
}

func NewPostRepository(db *sqlx.DB) *PostRepository {
	return &PostRepository{db: db}
}
