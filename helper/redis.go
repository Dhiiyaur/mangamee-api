package helper

import (
	"fmt"
	"mangamee-api/entity"
	"os"
)

func GenerateKey(params entity.RequestParams) string {
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

func GetExpiredTimeRedis() int {

	expEnv := os.Getenv("CACHE_TIME")
	var x int
	fmt.Sscanf(expEnv, "%d", &x)
	return x
}
