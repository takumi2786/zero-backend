package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/takumi2786/zero-backend/internal/domain"
	"github.com/takumi2786/zero-backend/internal/driver"
)

type UserRepository struct {
	db *sqlx.DB
}

func (r *UserRepository) FindUsers(ctx context.Context) (domain.Users, error) {
	queryer := driver.GetQueryer(r.db) // 読み取り専用
	sql := "SELECT u.id, u.name, u.email FROM users AS u"
	var users domain.Users
	if err := queryer.SelectContext(ctx, &users, sql); err != nil {
		return users, err
	}
	return users, nil
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}
