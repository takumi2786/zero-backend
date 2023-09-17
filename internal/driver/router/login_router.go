package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/takumi2786/zero-backend/internal/driver"
	"github.com/takumi2786/zero-backend/internal/interface/controller"
	"github.com/takumi2786/zero-backend/internal/interface/repository"
	"github.com/takumi2786/zero-backend/internal/usecase"
	"go.uber.org/zap"
)

func NewLoginRouter(
	config *driver.Config,
	timeout time.Duration,
	logger *zap.Logger,
	group *gin.RouterGroup,
	db *sqlx.DB,
) {
	ur := repository.NewUserRepository(db)
	ar := repository.NewAuthUserRepository(db)
	tg := driver.NewJWTTokenGenerator()
	lu := usecase.NewLoginInteractor(
		logger,
		ur,
		ar,
		tg,
	)
	lc := controller.LoginControler{
		LoginUsecase: lu,
		Logger:       logger,
	}
	group.POST("/login", lc.Login)
}
