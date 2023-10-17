package controller

import (
	"net/http"

	"github.com/takumi2786/zero-backend/internal/application/usecase"
	"go.uber.org/zap"
	"golang.org/x/xerrors"
)

type LoginController struct {
	Logger        *zap.Logger
	ILoginUsecase usecase.ILoginUsecase
}

func NewLoginController(logger *zap.Logger, lu usecase.ILoginUsecase) *LoginController {
	return &LoginController{ILoginUsecase: lu, Logger: logger}
}

type LoginRequestBody struct {
	Identifier string `json:"identifier" binding:"required"`
	Credential string `json:"credencial" binding:"required"`
}

type LoginResponseBody struct {
	Token string `json:"token"`
}

func (lc *LoginController) Login(loginRequest LoginRequestBody) (*SuccessResponse, *ErrorResponse) {
	token, err := lc.ILoginUsecase.Login("email", loginRequest.Identifier, loginRequest.Credential)
	if err != nil {
		if xerrors.As(err, &usecase.FailedToAuthorise) {
			return nil, &ErrorResponse{Code: http.StatusUnauthorized, Message: http.StatusText(http.StatusUnauthorized)}
		} else if xerrors.As(err, &usecase.FailedToGenerateToken) {
			return nil, &ErrorResponse{Code: http.StatusUnauthorized, Message: http.StatusText(http.StatusUnauthorized)}
		}
		lc.Logger.Error("failed to login", zap.Error(err))
		return nil, &ErrorResponse{Code: http.StatusInternalServerError, Message: http.StatusText(http.StatusInternalServerError)}
	}
	return &SuccessResponse{Code: http.StatusOK, Body: LoginResponseBody{Token: *token}}, nil
}
