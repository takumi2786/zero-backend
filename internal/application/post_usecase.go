package usecase

import (
	"context"
	"time"

	"github.com/takumi2786/zero-backend/internal/domain"
	"github.com/takumi2786/zero-backend/internal/infrastructure/repository"
)

type PostUseCase struct {
	PostRepository repository.PostRepository
	contextTimeout time.Duration
}

func NewPostUseCase(postRepository repository.PostRepository, contextTimeout time.Duration) *PostUseCase {
	return &PostUseCase{
		PostRepository: postRepository,
		contextTimeout: contextTimeout,
	}
}

func (pu *PostUseCase) Add(ctx context.Context, post *domain.Post) error {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()

	err := pu.PostRepository.AddPost(ctx, post)
	if err != nil {
		return err
	}
	return nil
}
