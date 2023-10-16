//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/takumi2786/zero-backend/internal/application/usecase"
	"github.com/takumi2786/zero-backend/internal/interface/controller"
	"github.com/takumi2786/zero-backend/internal/interface/repository"
	"github.com/takumi2786/zero-backend/internal/util"
	"go.uber.org/zap"
)

func InitializeLoginController(cfg *util.Config, logger *zap.Logger, db *sqlx.DB) *controller.LoginController {
	wire.Build(
		repository.NewUserRepository,
		repository.NewAuthUserRepository,
		usecase.NewJWTTokenGenerator,
		usecase.NewLoginUsecase,
		controller.NewLoginController,
	)
	return &controller.LoginController{}
}
