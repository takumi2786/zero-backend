package driver

import "github.com/takumi2786/zero-backend/internal/usecase"

type JWTTokenGenerator struct {
}

func NewJWTTokenGenerator() usecase.TokenGenerator {
	return &JWTTokenGenerator{}
}

func (jtg *JWTTokenGenerator) GenerateToken(id int64) (*string, error) {
	token := "token"
	return &token, nil
}
