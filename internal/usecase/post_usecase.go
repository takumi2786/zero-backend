package usecase

import "context"

type PostUsecase interface {
	AddPost(ctx context.Context, post AddPostInput) error
	FindPosts(ctx context.Context) (FindPostsOutput, error)
}
