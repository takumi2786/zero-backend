package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takumi2786/zero-backend/internal/usecase"
	"go.uber.org/zap"
	"golang.org/x/xerrors"
)

type LoginController struct {
	Logger       *zap.Logger
	LoginUsecase usecase.LoginUsecase
}

func NewLoginController(logger *zap.Logger, lu usecase.LoginUsecase) *LoginController {
	return &LoginController{LoginUsecase: lu, Logger: logger}
}

type LoginInput struct {
	Identifier string `json:"identifier" binding:"required"`
	Credential string `json:"credencial" binding:"required"`
}

func (lc *LoginController) Login(ctx *gin.Context) {
	// 入力値を取得
	var loginInput LoginInput // request body
	if err := ctx.ShouldBindJSON(&loginInput); err != nil {
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		lc.Logger.Error("failed to parse request", zap.Error(err))
		return
	}

	// usecaseのメソッドを呼び出す
	token, err := lc.LoginUsecase.Login(ctx, "email", loginInput.Identifier, loginInput.Credential)
	if err != nil {
		if xerrors.As(err, &usecase.FailedToAuthorise) {
			ctx.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			return
		} else if xerrors.As(err, &usecase.FailedToGenerateToken) {
			ctx.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			return
		}
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		lc.Logger.Error("failed to login", zap.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": *token})
}
