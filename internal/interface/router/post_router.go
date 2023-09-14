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

func NewPostRouter(config *driver.Config, timeout time.Duration, logger *zap.Logger, group *gin.RouterGroup, db *sqlx.DB) {
	pr := repository.NewPostRepository(db)
	pu := usecase.NewPostUseCase(pr, timeout)
	pc := controller.PostControler{
		PostUsecase: pu,
		Logger:      logger,
	}
	group.POST("/posts", pc.AddPost)
	group.GET("/posts", pc.FindPosts)
}
