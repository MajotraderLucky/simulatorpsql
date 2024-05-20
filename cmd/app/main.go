package main

import (
	"fmt"
	"log"
	"os"
	"simulatorpsql/internal/storage"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func loadEnv() (string, int, string, string, string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", 0, "", "", "", fmt.Errorf("error loading .env file: %v", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Convert port from string to int
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return "", 0, "", "", "", fmt.Errorf("invalid port value: %v", err)
	}

	return host, portInt, user, password, dbname, nil
}

func main() {
	host, port, user, password, dbname, err := loadEnv()
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}
	// Формируем строку подключения
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	dbHandler, err := storage.NewDBHandler(connectionString)
	if err != nil {
		log.Fatalf("Failed to initialize database handler: %v", err)
	}
	err = dbHandler.PingDatabase()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	fmt.Println("Successfully connected and pinged PostgreSQL database!")
}
