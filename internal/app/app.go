// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/awgst/datings/config"
	v1 "github.com/awgst/datings/internal/controller/http/v1"
	"github.com/awgst/datings/internal/usecase"
	"github.com/awgst/datings/pkg/app"
	"github.com/awgst/datings/pkg/database"
	"github.com/awgst/datings/pkg/httpserver"
	"github.com/awgst/datings/pkg/logger"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	errorLogger := logger.New()

	// Database
	db := database.New(database.ConnectionConfig{
		DbUrl: cfg.Database.URL,
	})

	// App
	app := &app.App{
		Config: cfg,
		DB:     db,
		Logger: errorLogger,
	}

	// Usecase
	uc := usecase.New(app)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, uc)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		errorLogger.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		errorLogger.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err := httpServer.Shutdown()
	if err != nil {
		errorLogger.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
