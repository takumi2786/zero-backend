package routes

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/takumi2786/zero-backend/internal/config"
	"github.com/takumi2786/zero-backend/internal/handlers"
)

// 認証が不要なルーティング
func SetRouting(ctx context.Context, router *gin.Engine, cfg *config.Config) error {

	// set handleers
	healthCheckhandler := handlers.NewHealthCheckHandler()
	router.GET("/healthcheck", healthCheckhandler.ServeHTTP)
	return nil
}

// // 認証を必要とするルーティング
// func SetAuthRouting(ctx context.Context, db *sqlx.DB, router *gin.Engine) error {

// 	return nil
// }
