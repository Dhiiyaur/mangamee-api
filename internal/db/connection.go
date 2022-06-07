package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func CreateConnection() *sql.DB {

	db, err := sql.Open("postgres", viper.GetString("POSTGRES_URL"))
	if err != nil {
		log.Println(err)
	}

	return db
}
