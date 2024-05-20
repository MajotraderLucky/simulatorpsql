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
	cfg, err := config.LoadEnv()
	if err != nil {
		logger.Log.Error("Failed to load .env file", zap.Error(err))
		os.Exit(1)
	}

	_, err = setupDatabase(cfg)
	if err != nil {
		logger.Log.Error("Failed to setup database", zap.Error(err))
		os.Exit(1)
	}

	logger.Log.Info("Successfully connected and pinged PostgreSQL database!")
}

func setupDatabase(cfg *config.Config) (*storage.DBHandler, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	dbHandler, err := storage.NewDBHandler(connectionString)
	if err != nil {
		return nil, err
	}

	if err = dbHandler.PingDatabase(); err != nil {
		return nil, err
	}

	return dbHandler, nil
}
