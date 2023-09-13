package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/takumi2786/zero-backend/internal/driver"
	"github.com/takumi2786/zero-backend/internal/infrastructure/repository"
	"github.com/takumi2786/zero-backend/internal/presentation/controller"
)

func NewUserRouter(config *driver.Config, group *gin.RouterGroup, db *sqlx.DB) {
	ur := repository.NewUserRepository(db)
	uc := controller.UserControler{
		UserRepository: *ur,
	}
	group.GET("/users", uc.GetUsers)
}
