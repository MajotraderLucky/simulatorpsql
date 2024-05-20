// config/config.go

package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv() (string, int, string, string, string, error) {
	if err := godotenv.Load(); err != nil {
		return "", 0, "", "", "", fmt.Errorf("error loading .env file: %v", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	portInt, err := strconv.Atoi(port)
	if err != nil {
		return "", 0, "", "", "", fmt.Errorf("invalid port value: %v", err)
	}

	return host, portInt, user, password, dbname, nil
}
