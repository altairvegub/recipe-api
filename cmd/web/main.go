package main

import (
	"context"
	"os"
	"os/signal"
	"recipe/internal/config"
	"recipe/internal/service"
	"syscall"
	"time"

	"recipe/internal/ports"
	"recipe/pkg/log"

	"recipe/internal/database"
	"recipe/internal/database/user"

	"go.uber.org/zap"
)

func main() {
	cfg := config.LoadConfigs(config.RecipeApiPrefix)
	logger := log.New()

	db, err := database.New(cfg.PostgresConfig.Port, cfg.PostgresConfig.Host, cfg.PostgresConfig.Username, cfg.PostgresConfig.Password, cfg.PostgresConfig.Database)
	if err != nil {
		logger.Errorf("Failed to initialize postgres", zap.Error(err))
	}

	svc := service.New(
		user.NewRepository(db),
	)
	srv := ports.RunHTTPServer(cfg.HTTPServer.Port, logger, svc)

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Errorf("Server forced to shutdown: %v", err)
	}

	logger.Info("Server exiting")
}
