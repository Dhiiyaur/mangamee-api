package db

import (
	"log"
	"time"
)

func InsertDataUserLog(Api, mangaSource, mangaTitle, mangaChapter, ip string) {

	db := CreateConnection()

	defer db.Close()

	dateNow := time.Now().UTC()
	_, err := db.Exec(`INSERT INTO logs (api_source, src, manga_title, manga_chapter, created_on, ip) VALUES ($1, $2, $3, $4, $5, $6)`, Api, mangaSource, mangaTitle, mangaChapter, dateNow, ip)
	if err != nil {
		log.Println(err)
	}
}
