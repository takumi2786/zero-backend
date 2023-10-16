package main

import (
	"context"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/takumi2786/zero-backend/internal/infrastructure"
	"github.com/takumi2786/zero-backend/internal/infrastructure/waf/gin"
	"github.com/takumi2786/zero-backend/internal/util"

	"go.uber.org/zap"
)

func main() {
	log.Printf("start server")
	appContext := context.Background()
	if err := run(appContext); err != nil {
		log.Printf("failed to terminated server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	// Read environment variables
	cfg, err := util.NewConfig()
	if err != nil {
		return err
	}

	// Setup logger
	zapCfg := zap.NewProductionConfig()
	zapCfg.DisableStacktrace = true
	logger, _ := zapCfg.Build()
	defer logger.Sync()

	// Connect to database
	db, err := infrastructure.NewDB(ctx, cfg)
	if err != nil {
		logger.Error("Failed to connect to database")
		panic(err)
	}
	ginApp := gin.NewGinApp(cfg, logger, db)
	lc := InitializeLoginController(cfg, logger, db)
	err = ginApp.Run(lc)
	if err != nil {
		panic(err)
	}

	return nil
}
