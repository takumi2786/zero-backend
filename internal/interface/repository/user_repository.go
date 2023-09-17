package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/takumi2786/zero-backend/internal/domain"
)

/*
User
*/
type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByUserId(userId int64) (*domain.User, error) {
	var user domain.User
	sql := "SELECT * FROM users WHERE user_id = ?"
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

func NewAuthUserRepository(db *sqlx.DB) *AuthUserRepository {
	return &AuthUserRepository{db: db}
}

func (r *AuthUserRepository) GetByIdentity(userId int64, identityType string, identity string) (*domain.UserAuth, error) {
	var authUser domain.UserAuth
	sql := "SELECT * FROM auth_users WHERE user_id = ? AND identity_type = ?"
	err := r.db.Get(&authUser, sql, userId, identityType)
	if err != nil {
		return nil, err
	}
	return &authUser, nil
}
