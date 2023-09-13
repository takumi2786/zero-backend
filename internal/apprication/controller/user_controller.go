package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takumi2786/zero-backend/internal/domain"
	"github.com/takumi2786/zero-backend/internal/infrastructure/repository"
)

type GetUsersResponse struct {
	Id    int    `json:"message"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserControler struct {
	UserRepository repository.UserRepository
}

func (gu *UserControler) GetUsers(ctx *gin.Context) {
	resp, err := gu.UserRepository.FindUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
