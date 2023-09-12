package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/takumi2786/zero-backend/internal/config"
	"github.com/takumi2786/zero-backend/internal/routes"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	srv *http.Server
}

func NewServer(mux http.Handler, addr string) *Server {
	return &Server{
		srv: &http.Server{Handler: mux, Addr: addr},
	}
}

func (s *Server) Run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		// http.ErrServerClosed は
		// http.Server.Shutdown() が正常に終了したことを示すので異常ではない
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})

	// チャンネルからの終了通知を待つ
	<-ctx.Done()
	// サーバーをシャットダウン
	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}
	// 別ルーチンのグレースフルシャットダウンの終了を待つ
	return eg.Wait()
}

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
	cfg, err := config.New()
	if err != nil {
		return err
	}
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	// Routing
	if err = routes.SetRouting(ctx, router, cfg); err != nil {
		return err
	}

	server := NewServer(router, fmt.Sprintf(":%d", cfg.Port))
	return server.Run(ctx)
}
