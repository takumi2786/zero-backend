package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/takumi2786/zero-backend/internal/driver"
	"github.com/takumi2786/zero-backend/internal/infrastructure/repository"
	"github.com/takumi2786/zero-backend/internal/presentation/controller"
)

func NewPostRouter(config *driver.Config, group *gin.RouterGroup, db *sqlx.DB) {
	pr := repository.NewPostRepository(db)
	pc := controller.PostControler{
		PostRepository: *pr,
	}
	group.POST("/posts", pc.AddPost)
}
