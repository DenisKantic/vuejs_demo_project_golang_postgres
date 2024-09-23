package db_connection

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DB_CONNECT string

func init() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		log.Fatal("Error loading .env file")
	}

	DB_CONNECT = os.Getenv("DATABASE_URL")

}

func DBConnect() (*sql.DB, error) {

	db, err := sql.Open("postgres", DB_CONNECT)
	if err != nil {
		return nil, fmt.Errorf("error oppening connection %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		// fmt.Println("Erorr connecting to database")
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	fmt.Println("CONNECTED TO THE DATABASE")
	return db, nil
}

func init() {
	_, err := DBConnect()
	if err != nil {
		return
	}
}
