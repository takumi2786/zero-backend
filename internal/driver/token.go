package driver

type JWTTokenGenerator struct {
}

func NewJWTTokenGenerator() *JWTTokenGenerator {
	return &JWTTokenGenerator{}
}

func (jtg *JWTTokenGenerator) GenerateToken(id int64) (*string, error) {
	token := "token"
	return &token, nil
}
