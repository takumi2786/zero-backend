package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/takumi2786/zero-backend/internal/driver"
	"go.uber.org/zap"
)

func Setup(config *driver.Config, logger *zap.Logger, gin *gin.Engine, db *sqlx.DB, timeout time.Duration) {
	publicRouter := gin.Group("")
	NewPostRouter(config, timeout, logger, publicRouter, db)
	NewLoginRouter(config, timeout, logger, publicRouter, db)
}
