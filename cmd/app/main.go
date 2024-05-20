package main

import (
	"os"
	"simulatorpsql/internal/config"
	"simulatorpsql/internal/logger"
	"simulatorpsql/internal/storage"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	logger.InitLogger()
	cfg, err := config.LoadEnv()
	if err != nil {
		logger.Log.Error("Failed to load .env file", zap.Error(err))
		os.Exit(1)
	}

	_, err = storage.SetupDatabase(cfg) // Или storage.SetupDatabase(cfg), в зависимости от выбранного местоположения
	if err != nil {
		logger.Log.Error("Failed to setup database", zap.Error(err))
		os.Exit(1)
	}

	logger.Log.Info("Successfully connected and pinged PostgreSQL database!")
}
