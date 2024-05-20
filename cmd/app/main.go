package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
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

func pingDatabase() error {
	host, port, user, password, dbname, err := loadEnv()
	if err != nil {
		return fmt.Errorf("failed to load env: %v", err)
	}

	// Формируем строку подключения
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Пытаемся подключиться к базе данных
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
	defer db.Close()

	// Пингуем базу данных
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error pinging database: %v", err)
	}

	fmt.Println("Successfully connected and pinged PostgreSQL database!")
	return nil
}

func main() {
	err := pingDatabase()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
}
