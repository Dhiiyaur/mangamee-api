package helper

import (
	"encoding/json"
	"fmt"
	"mangamee-api/entity"
)

func RepositoryToArrayResponseData(repo entity.MangaRepository) []entity.MangaData {

	var mangaData []entity.MangaData
	val, _ := json.Marshal(repo.Data)
	err := json.Unmarshal([]byte(val), &mangaData)
	if err != nil {
		fmt.Println("what")
	}
	return mangaData
}

func RepositoryToResponseData(repo entity.MangaRepository) entity.MangaData {

	var mangaData entity.MangaData
	val, _ := json.Marshal(repo.Data)
	err := json.Unmarshal([]byte(val), &mangaData)
	if err != nil {
		fmt.Println("what")
	}
	return mangaData
}
