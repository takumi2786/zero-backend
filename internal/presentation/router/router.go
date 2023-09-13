package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/takumi2786/zero-backend/internal/driver"
)

func Setup(config *driver.Config, gin *gin.Engine, db *sqlx.DB) {
	publicRouter := gin.Group("")
	NewUserRouter(config, publicRouter, db)
	NewPostRouter(config, publicRouter, db)
}
