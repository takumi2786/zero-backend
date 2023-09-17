package usecase

type TokenGenerator interface {
	GenerateToken(userId int64) (*string, error)
}
