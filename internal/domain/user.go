package domain

type User struct {
	UserId int64  `json:"userId" db:"user_id"`
	Name   string `json:"name" db:"name"`
}

type UserAuth struct {
	Id           int64  `json:"id" db:"id"`
	UserId       int64  `json:"userId" db:"user_id"`
	IdentityType string `json:"identityType" db:"identity_type"`
	Identifier   string `json:"identifier" db:"identifier"`
	Credential   string `json:"credential" db:"credential"`
}

type UserRepository interface {
	GetByUserId(userId int64) (*User, error)
}

type AuthUserRepository interface {
	GetByIdentifier(identityType string, identifier string) (*UserAuth, error)
}
