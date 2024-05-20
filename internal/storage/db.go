package storage

import (
	"database/sql"
	"fmt"
	"simulatorpsql/internal/config"
	"simulatorpsql/internal/logger"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type DBHandler struct {
	db *sql.DB
}

func NewDBHandler(connectionString string) (*DBHandler, error) {
	// Добавляем sslmode=disable для отключения SSL
	connectionString += " sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		logger.Log.Error("Failed to open database connection", zap.String("connectionString", connectionString), zap.Error(err))
		return nil, err
	}
	return &DBHandler{db: db}, nil
}

func (handler *DBHandler) PingDatabase() error {
	err := handler.db.Ping()
	if err != nil {
		// Правильно логируем ошибку перед возвратом её
		logger.Log.Error("Failed to ping database", zap.Error(err))
		return err
	}
	return nil
}

// В storage.go
func SetupDatabase(cfg *config.Config) (*DBHandler, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	return NewDBHandler(connectionString)
}
