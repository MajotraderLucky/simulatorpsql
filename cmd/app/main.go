package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"        // Или используйте имя хоста Docker, если запускаете из другого контейнера
	port     = 5432               // Порт, на котором работает PostgreSQL
	user     = "your_db_user"     // Имя пользователя, указанное в docker-compose.yml
	password = "your_db_password" // Пароль пользователя, указанный в docker-compose.yml
	dbname   = "your_db_name"     // Имя базы данных, указанное в docker-compose.yml
)

func pingDatabase() error {
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
