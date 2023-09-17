package usecase

type TokenGenerator interface {
	GenerateToken(id int64) (*string, error)
}
