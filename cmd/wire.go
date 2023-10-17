//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/takumi2786/zero-backend/internal/application/usecase"
	"github.com/takumi2786/zero-backend/internal/interfaces/controller"
	"github.com/takumi2786/zero-backend/internal/interfaces/repository"
	"github.com/takumi2786/zero-backend/internal/util"
	"go.uber.org/zap"
)

func InitializeLoginController(cfg *util.Config, logger *zap.Logger, sqlHandler repository.SQLHandler) *controller.LoginController {
	wire.Build(
		repository.NewUserRepository,
		repository.NewAuthUserRepository,
		usecase.NewJWTTokenGenerator,
		usecase.NewLoginUsecase,
		controller.NewLoginController,
	)
	return &controller.LoginController{}
}
