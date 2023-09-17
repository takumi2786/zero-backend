package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/takumi2786/zero-backend/internal/domain"
)

/*
User
*/
type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) domain.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByUserId(ctx context.Context, userId int64) (*domain.User, error) {
	var user domain.User
	sql := "SELECT user_id, name FROM users WHERE user_id = ?"
	err := r.db.Get(&user, sql, userId)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

/*
AuthUser
*/
type AuthUserRepository struct {
	db *sqlx.DB
}

func NewAuthUserRepository(db *sqlx.DB) domain.AuthUserRepository {
	return &AuthUserRepository{db: db}
}

func (r *AuthUserRepository) GetByIdentifier(ctx context.Context, identityType string, identifier string) (*domain.UserAuth, error) {
	var authUser domain.UserAuth
	sql := "SELECT * FROM user_auths WHERE identifier = ? AND identity_type = ?"
	err := r.db.Get(&authUser, sql, identifier, identityType)
	if err != nil {
		return nil, err
	}
	return &authUser, nil
}
