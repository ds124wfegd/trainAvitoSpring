package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ds124wfegd/trainAvitoSpring/config"
	"github.com/ds124wfegd/trainAvitoSpring/internal/httpServer"
	"go.uber.org/zap"
)

func main() {
	Logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to initialize zap logger: %v", err)
	}
	defer Logger.Sync()

	viperInstance, err := config.LoadConfig()

	if err != nil {
		Logger.Error("Cannot load config.",
			zap.Error(err))
	}

	cfg, err := config.ParseConfig(viperInstance)
	if err != nil {
		Logger.Error("Cannot parse config",
			zap.Error(err))
	}

	server := httpServer.NewServer(cfg)
	if err = server.Run(); err != nil {
		Logger.Error("",
			zap.Error(err))
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	Logger.Info("app Shutting Down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		Logger.Error("Error occured on server shutting down",
			zap.Error(err))
	}

}
