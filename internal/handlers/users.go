package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetUsersResponse struct {
	Id       int    `json:"message"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	CreateAt string `json:"create_at"`
}

type GetUsers struct{}

func (gu *GetUsers) ServeHTTP(ctx *gin.Context) {
	resp := GetUsersResponse{
		Id:       1,
		Name:     "takumi",
		Email:    "mail@example.com",
		CreateAt: "2020-01-01",
	}
	ctx.JSON(http.StatusOK, resp)
}

func NewGetUsersHandler() *GetUsers {
	return &GetUsers{}
}
