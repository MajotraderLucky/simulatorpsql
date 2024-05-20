package main

import (
	"fmt"
	"log"
	"simulatorpsql/internal/config"
	"simulatorpsql/internal/storage"

	_ "github.com/lib/pq"
)

func main() {
	host, port, user, password, dbname, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dbHandler, err := storage.NewDBHandler(connectionString)
	if err != nil {
		log.Fatalf("Failed to initialize database handler: %v", err)
	}
	if err = dbHandler.PingDatabase(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	fmt.Println("Successfully connected and pinged PostgreSQL database!")
}
