package config

import (
	"errors"
	"os"
	"strconv"

	"simulatorpsql/internal/logger"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

// Config структура для хранения конфигурации базы данных.
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// LoadEnv загружает параметры конфигурации из .env файла и возвращает их в виде структуры Config.
func LoadEnv() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		logger.Log.Error("Error loading .env file", zap.Error(err))
		return nil, errors.New("failed to load environment variables") // Возвращаем обобщенное сообщение об ошибке
	}

	// Получение переменных окружения
	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Преобразование порта из строки в число
	port, err := strconv.Atoi(portStr)
	if err != nil {
		logger.Log.Error("Invalid port value", zap.String("port", portStr), zap.Error(err))
		return nil, errors.New("invalid port value") // Возвращаем обобщенное сообщение об ошибке
	}

	// Возвращение структуры Config с загруженными данными
	return &Config{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbname,
	}, nil
}
