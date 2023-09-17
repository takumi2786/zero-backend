package usecase

import (
	"context"
)

type LoginUsecase interface {
	Login(ctx context.Context, identityType string, identifier string, credential string) (*string, error)
}
