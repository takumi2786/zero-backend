package usecase

import "context"

type PostUseCase interface {
	AddPost(ctx context.Context, post AddPostInput) error
	FindPosts(ctx context.Context) (FindPostsOutput, error)
}
