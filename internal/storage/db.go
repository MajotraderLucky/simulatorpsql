package storage

import (
	"database/sql"
	"simulatorpsql/internal/logger"

	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

type DBHandler struct {
	db *sql.DB
}

func NewDBHandler(connectionString string) (*DBHandler, error) {
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
