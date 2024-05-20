package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBHandler struct {
	db *sql.DB
}

func NewDBHandler(connectionString string) (*DBHandler, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return &DBHandler{db: db}, nil
}

func (handler *DBHandler) PingDatabase() error {
	err := handler.db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}
	return nil
}
