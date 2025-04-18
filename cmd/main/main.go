package main

import (
	"log"

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
}
