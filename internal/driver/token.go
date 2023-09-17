package driver

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/takumi2786/zero-backend/internal/usecase"
)

type JWTTokenGenerator struct {
}

func NewJWTTokenGenerator() usecase.TokenGenerator {
	return &JWTTokenGenerator{}
}

func (jtg *JWTTokenGenerator) GenerateToken(userId int64) (*string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(""))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}
