package helper

import (
	"encoding/json"
	"fmt"
	"mangamee-api/src/entity"
)

func GenerateKey(params entity.MangaParams) string {
	var redisKey string
	switch params.Path {
	case "index":
		redisKey = fmt.Sprintf("%s-%s-%s", params.Path, params.Source, params.PageNumber)
	case "detail":
		redisKey = fmt.Sprintf("%s-%s-%s", params.Path, params.Source, params.MangaId)
	case "search":
		redisKey = fmt.Sprintf("%s-%s-%s", params.Path, params.Source, params.Search)
	case "read":
		redisKey = fmt.Sprintf("%s-%s-%s-%s", params.Path, params.Source, params.MangaId, params.ChapterId)
	case "chapter":
		redisKey = fmt.Sprintf("%s-%s-%s", params.Path, params.Source, params.MangaId)
	}
	return redisKey
}

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
