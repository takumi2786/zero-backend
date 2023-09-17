package usecase

import (
	"context"
	"time"

	"github.com/takumi2786/zero-backend/internal/domain"
)

type PostInteractor struct {
	PostRepository domain.PostRepository
	contextTimeout time.Duration
}

func NewPostInteractor(postRepository domain.PostRepository, contextTimeout time.Duration) PostUsecase {
	return &PostInteractor{
		PostRepository: postRepository,
		contextTimeout: contextTimeout,
	}
}

type AddPostInput struct {
	Title   string
	Content string
}

func (pu *PostInteractor) AddPost(ctx context.Context, post AddPostInput) error {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()

	postInDB := domain.Post{
		Title:     post.Title,
		Content:   post.Content,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	err := pu.PostRepository.AddPost(ctx, postInDB)
	if err != nil {
		return err
	}
	return nil
}

type PostElement struct {
	Id      domain.PostID `json:"id"`
	Title   string        `json:"title"`
	Content string        `json:"content"`
}
type FindPostsOutput []PostElement

func (pu *PostInteractor) FindPosts(ctx context.Context) (FindPostsOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()

	posts, err := pu.PostRepository.FindPosts(ctx)
	if err != nil {
		return nil, err
	}

	var response = make(FindPostsOutput, len(posts))
	for _, post := range posts {
		response = append(
			response,
			PostElement{
				Id:      post.Id,
				Title:   post.Title,
				Content: post.Content,
			},
		)
	}
	return response, nil
}
