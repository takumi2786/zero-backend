package domain

import "context"

type User struct {
	UserId int64  `json:"userId" db:"user_id"`
	Name   string `json:"name" db:"name"`
}

type UserAuth struct {
	Id           int64  `json:"id" db:"id"`
	UserId       int64  `json:"userId" db:"user_id"`
	IdentityType string `json:"identityType" db:"identity_type"`
	Identity     string `json:"identity" db:"identity"`
	Credential   string `json:"credential" db:"credential"`
}

type UserRepository interface {
	GetByUserId(ctx context.Context, userId int64) (*User, error)
}

type AuthUserRepository interface {
	GetByIdentity(ctx context.Context, identityType string, identity string) (*UserAuth, error)
}
