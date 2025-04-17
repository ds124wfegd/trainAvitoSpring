package main

import (
	"log"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction
	if err!= nil {
		log.Fatalf("Failed to initialize zap logger: %v", err)
	}
	defer logger.Sync()

	viperInstance, err := config.LoadConfig()

	if err != nil {
		logger.Error("Cannot load config.", 
		zap.Error(err))
	}

	cfg, err := config.ParseConfig(viperInstance)
	if err!=nil {
		logger.Error("Cannot parse config", 
		zap.Error(err))
	}

	server:=httpServer.NewServer(cfg)
	if err = server.Run() != nil {
		logger.Error("Cannot parse config", 
		zap.Error(err))
	}
}