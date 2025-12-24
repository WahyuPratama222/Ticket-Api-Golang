package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() error {
	// Load environment
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open DB: %v", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping DB: %v", err)
	}

	DB = db
	fmt.Println("Connected to MySQL database successfully")
	return nil
}
