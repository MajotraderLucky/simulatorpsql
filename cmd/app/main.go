package main

import (
	"fmt"
	"os"
	"simulatorpsql/internal/config"
	"simulatorpsql/internal/logger"
	"simulatorpsql/internal/storage"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	logger.InitLogger()
	logger.Log.Info("Starting the simulator for PostgreSQL database...")
	cfg, err := config.LoadEnv()
	if err != nil {
		logger.Log.Error("Failed to load environment variables", zap.Error(err))
		os.Exit(1)
	}

	// Используйте поля из cfg для формирования строки подключения
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	dbHandler, err := storage.NewDBHandler(connectionString)
	if err != nil {
		logger.Log.Error("Failed to initialize database handler", zap.Error(err))
		os.Exit(1)
	}

	if err = dbHandler.PingDatabase(); err != nil {
		logger.Log.Error("Failed to ping database", zap.Error(err))
		os.Exit(1)
	}

	logger.Log.Info("Successfully connected and pinged PostgreSQL database!")
}
