package repository

import (
	"github.com/takumi2786/zero-backend/internal/domain"
)

/*
ここには、SQLHanndlerを利用したデータベースへのアクセス処理を実装する。
*/

/*
User
*/
type UserRepository struct {
	sqlHandler SQLHandler
}

func NewUserRepository(sqlHandler SQLHandler) domain.UserRepository {
	return &UserRepository{sqlHandler: sqlHandler}
}

func (r *UserRepository) GetByUserId(userId int64) (*domain.User, error) {
	var user domain.User
	sql := "SELECT user_id, name FROM users WHERE user_id = ?"
	err := r.sqlHandler.Get(&user, sql, userId)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

/*
AuthUser
*/
type AuthUserRepository struct {
	sqlHandler SQLHandler
}

func NewAuthUserRepository(sqlHandler SQLHandler) domain.AuthUserRepository {
	return &AuthUserRepository{sqlHandler: sqlHandler}
}

func (r *AuthUserRepository) GetByIdentifier(identityType string, identifier string) (*domain.UserAuth, error) {
	var authUser domain.UserAuth
	sql := "SELECT * FROM user_auths WHERE identifier = ? AND identity_type = ?"
	err := r.sqlHandler.Get(&authUser, sql, identifier, identityType)
	if err != nil {
		return nil, err
	}
	return &authUser, nil
}
