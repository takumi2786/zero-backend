package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/takumi2786/zero-backend/internal/driver"
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
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	gin := gin.Default()

	// Setup logger
	zapCfg := zap.NewProductionConfig()
	zapCfg.DisableStacktrace = true
	logger, _ := zapCfg.Build()
	defer logger.Sync()

	logger.Info("Start Server")

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	gin.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	// gin.Use(ginzap.RecoveryWithZap(logger, true))

	// Connect to database
	db, err := driver.NewDB(ctx, cfg)
	if err != nil {
		panic(err)
	}

	// CORS
	gin.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	/*
		Routing
	*/

	// Setup dependencies
	timeout := time.Duration(cfg.Timeout) * time.Second
	lc := InitializeLoginController(cfg, logger, db)
	pc := InitializePostController(logger, db, timeout)

	// setup router
	driver.SetRouting(
		gin,
		lc,
		pc,
	)

	// Run server
	err = gin.Run(fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		panic(err)
	}
	return nil
}
