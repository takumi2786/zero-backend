package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takumi2786/zero-backend/internal/application/usecase"
	"github.com/takumi2786/zero-backend/internal/domain"
	"github.com/takumi2786/zero-backend/internal/util"
	"go.uber.org/zap"
)

type PostController struct {
	config       *util.Config
	logger       *zap.Logger
	IPostUsecase usecase.IPostUsecase
}

func NewPostController(cfg *util.Config, logger *zap.Logger, pu usecase.IPostUsecase) *PostController {
	return &PostController{config: cfg, IPostUsecase: pu, logger: logger}
}

/*
作成系
*/
type AddPostInput struct {
	Title   string `json:"title" binding:"required,max=50"`
	Content string `json:"content" binding:"required,max=500"`
}

func (pc *PostController) AddPost(ctx *gin.Context) {
	// 入力値を取得
	var postInput AddPostInput // request body
	if err := ctx.ShouldBindJSON(&postInput); err != nil {
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		pc.logger.Error("failed to parse request", zap.Error(err))
		return
	}

	// usecaseのメソッドを呼び出す
	err := pc.IPostUsecase.AddPost(
		ctx, usecase.AddPostInput{Title: postInput.Title, Content: postInput.Content},
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		pc.logger.Error("failed to add post", zap.Error(err))
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "created"})
}

/*
取得系
*/
type PostElement struct {
	Id      domain.PostID `json:"id"`
	Title   string        `json:"title"`
	Content string        `json:"content"`
}
type FindPostsOutPut []PostElement

func (pc *PostController) FindPosts(ctx *gin.Context) {
	posts, err := pc.IPostUsecase.FindPosts(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		pc.logger.Error("failed to find posts", zap.Error(err))
		return
	}

	results := make(FindPostsOutPut, len(posts))
	for index, post := range posts {
		results[index] = PostElement{Id: post.Id, Title: post.Title, Content: post.Content}
	}
	ctx.JSON(http.StatusOK, results)
}
