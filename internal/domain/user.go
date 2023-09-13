package domain

import (
	"context"
	"time"
)

// model
type UserID int64

type User struct {
	ID        UserID    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type Users []*User

// repository interface
type UserRepository interface {
	FindUsers(ctx context.Context) (Users, error)
}
