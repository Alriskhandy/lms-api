package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/myorg/lms-backend/internal/config"
	"github.com/myorg/lms-backend/internal/http/middleware"
	"github.com/myorg/lms-backend/internal/logger"
)

func main() {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	// Init logger
	if err := logger.Init(cfg.Env == "production"); err != nil {
		panic("failed to init logger: " + err.Error())
	}
	defer logger.Sync()

	// Setup Gin router
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestID())
	r.Use(middleware.RequestLogger())

	// Healthcheck
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Setup HTTP server
	srv := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: r,
	}

	// Run server in goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Sugar().Fatalf("listen: %s\n", err)
		}
	}()
	logger.Log.Sugar().Infof("server started on %s", cfg.AppPort)

	// Graceful shutdown dengan signal.NotifyContext
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-ctx.Done() // tunggu signal Ctrl+C
	logger.Log.Sugar().Info("shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Log.Sugar().Fatal("server forced to shutdown:", err)
	}

	logger.Log.Sugar().Info("server exiting")
}
