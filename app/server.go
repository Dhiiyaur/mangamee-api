package app

import (
	"log"
	"mangamee-api/logger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"

	MangaController "mangamee-api/controller/manga"
	MangaRepository "mangamee-api/repository/manga"
	MangaService "mangamee-api/service/manga"

	BookmarkController "mangamee-api/controller/bookmark"
	BookmarkRepository "mangamee-api/repository/bookmark"
	BookmarkService "mangamee-api/service/bookmark"

	LinkController "mangamee-api/controller/link"
	LinkRepository "mangamee-api/repository/link"
	LinkService "mangamee-api/service/link"
)

func Run() {

	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", zap.Error(err))
		log.Fatal("Error loading .env file")
	}

	logger.InitLogger()
	ds, err := InitDS()
	if err != nil {
		logger.Error("Unable to initialize data source", zap.Error(err))
		log.Fatalln("Unable to initialize data source, %w", err)
	}

	srv := gin.New()
	MiddlewareSetup(srv)

	mangaRoute := srv.Group("/manga")
	MangaRepository := MangaRepository.NewMangaRepository(ds.RedisClient, ds.DB)
	MangaService := MangaService.NewMangaService(MangaRepository)
	MangaController := MangaController.NewMangaController(MangaService)
	MangaController.Mount(mangaRoute)

	bookmarkRoute := srv.Group("/bookmark")
	BookmarkRepository := BookmarkRepository.NewBookmarkRepository(ds.RedisClient)
	BookmarkService := BookmarkService.NewBookmarkService(BookmarkRepository)
	BookmarkController := BookmarkController.NewBookmarkController(BookmarkService)
	BookmarkController.Mount(bookmarkRoute)

	linkRoute := srv.Group("/link")
	LinkRepository := LinkRepository.NewLinkRepository(ds.RedisClient)
	LinkService := LinkService.NewLinkService(LinkRepository)
	LinkController := LinkController.NewLinkController(LinkService)
	LinkController.Mount(linkRoute)

	srv.Run()
}
