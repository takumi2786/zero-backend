//go:build wireinject
// +build wireinject

package main

import (
	"time"

	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/takumi2786/zero-backend/internal/driver"
	"github.com/takumi2786/zero-backend/internal/interface/controller"
	"github.com/takumi2786/zero-backend/internal/interface/repository"
	"github.com/takumi2786/zero-backend/internal/usecase"
	"go.uber.org/zap"
)

func InitializeLoginController(logger *zap.Logger, db *sqlx.DB) *controller.LoginController {
	wire.Build(
		repository.NewUserRepository,
		repository.NewAuthUserRepository,
		driver.NewJWTTokenGenerator,
		usecase.NewLoginInteractor,
		controller.NewLoginController,
	)
	return &controller.LoginController{}
}

func InitializePostController(logger *zap.Logger, db *sqlx.DB, ontextTimeout time.Duration) *controller.PostController {
	wire.Build(
		repository.NewPostRepository,
		usecase.NewPostInteractor,
		controller.NewPostController,
	)
	return &controller.PostController{}
}
