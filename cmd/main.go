package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/takumi2786/zero-backend/internal/driver"
	"github.com/takumi2786/zero-backend/internal/presentation/router"
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
	cfg, err := driver.NewConfig()
	if err != nil {
		return err
	}
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Connect to database
	db, err := driver.NewDB(ctx, cfg)
	if err != nil {
		panic(err)
	}
	gin := gin.Default()

	// CORS
	gin.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	// Routing
	router.Setup(cfg, gin, db)

	// Run server
	err = gin.Run(fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		panic(err)
	}
	return nil
}
