package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func CreateConnection() *sql.DB {

	err := godotenv.Load(".env")

	if err != nil {
		log.Println(err)
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		log.Println(err)
	}

	return db
}
