package app

import (
	"mangamee-api/src/config"
	MangaController "mangamee-api/src/controller/manga"
	MangaRepository "mangamee-api/src/repository/manga"
	MangaService "mangamee-api/src/service/manga"

	ShortenerController "mangamee-api/src/controller/shortener"
	ShortenerRepository "mangamee-api/src/repository/shortener"
	ShortenerService "mangamee-api/src/service/shortener"

	BookmarkController "mangamee-api/src/controller/bookmark"
	BookmarkRepository "mangamee-api/src/repository/bookmark"
	BookmarkService "mangamee-api/src/service/bookmark"

	"net/http"

	"github.com/labstack/echo/v4"
)

func Run() {

	config.ReadConfig()

	e := echo.New()
	MiddlewareSetup(e)

	db, err := CreateDbConnection(config.Cfg)
	if err != nil {
		config.Logger.Error(err)
	}

	redis, err := CreateRedisConnection(config.Cfg)
	if err != nil {
		config.Logger.Error(err)
	}

	mangaRoute := e.Group("/manga")
	MangaRepository := MangaRepository.NewMangaRepository(redis, db)
	MangaService := MangaService.NewMangaService(MangaRepository)
	MangaController := MangaController.NewMangaController(MangaService)
	MangaController.Mount(mangaRoute)

	shortenerRoute := e.Group("/link")
	ShortenerRepository := ShortenerRepository.NewShortenerRepository(redis)
	ShortenerService := ShortenerService.NewShortenerService(ShortenerRepository)
	ShortenerController := ShortenerController.NewShortenerController(ShortenerService)
	ShortenerController.Mount(shortenerRoute)

	bookmarkRoute := e.Group("/bookmark")
	BookmarkRepository := BookmarkRepository.NewBookmarkRepository(redis)
	BookmarkService := BookmarkService.NewBookmarkService(BookmarkRepository)
	BookmarkController := BookmarkController.NewBookmarkController(BookmarkService)
	BookmarkController.Mount(bookmarkRoute)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome To MangameeApi")
	})

	config.Logger.Info("server start")
	e.Start(":" + config.Cfg.Server.Port)
}
