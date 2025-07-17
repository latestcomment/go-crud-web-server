package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	ssl_mode := os.Getenv("SSL_MODE")

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		db_user, db_pass, db_name, db_host, db_port, ssl_mode,
	)

	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("DB open error:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("DB ping error:", err)
	}

	log.Println("Connected to PostgreSQL.")
}
