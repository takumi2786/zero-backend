package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takumi2786/zero-backend/internal/infrastructure/repository"
)

type PostControler struct {
	PostRepository repository.PostRepository
}

func (pc *PostControler) AddPost(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}
