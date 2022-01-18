package main

import (
	"context"
	"os"
	"os/signal"
	"recipe/internal/config"
	"syscall"
	"time"

	"recipe/internal/ports"
	"recipe/pkg/log"
)

func main() {
	cfg := config.LoadConfigs(config.RecipeApiPrefix)
	logger := log.New()

	srv := ports.RunHTTPServer(cfg.HTTPServer.Port, logger)

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
