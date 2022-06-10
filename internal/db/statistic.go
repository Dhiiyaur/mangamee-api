package db

import (
	"log"
)

func InsertDataUserLog(Api string, mangaSource int, mangaTitle string, mangaChapter string) {

	db := CreateConnection()

	defer db.Close()

	// dateNow := time.Now().UTC()
	// _, err := db.Exec(`INSERT INTO logs (api_source, src, manga_title, manga_chapter, created_on, ip) VALUES ($1, $2, $3, $4, $5, $6)`, Api, mangaSource, mangaTitle, mangaChapter, dateNow, ip)

	_, err := db.Exec(`INSERT INTO logs (api, source, title, chapter) VALUES ($1, $2, $3, $4)`, Api, mangaSource, mangaTitle, mangaChapter)
	if err != nil {
		log.Println(err)
	}
}
