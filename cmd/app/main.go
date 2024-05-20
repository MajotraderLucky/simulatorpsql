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
	host, port, user, password, dbname, err := config.LoadEnv()
	if err != nil {
		logger.Log.Error("Failed to load environment variables", zap.Error(err))
		os.Exit(1)
	}
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
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
