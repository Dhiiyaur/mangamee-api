package app

import (
	"mangamee-api/internal/config"
	bookmarkcontroller "mangamee-api/internal/controller/bookmark"
	shortenercontroller "mangamee-api/internal/controller/link_shortener"
	mangacontroller "mangamee-api/internal/controller/manga"

	"mangamee-api/internal/repository"
	bookmarkservice "mangamee-api/internal/service/bookmark"
	mangaservice "mangamee-api/internal/service/manga"
	linkshortenerservice "mangamee-api/internal/service/shortener"

	"net/http"

	"github.com/labstack/echo/v4"
)

func Run() {

	config.ReadConfig()

	e := echo.New()
	MiddlewareSetup(e)

	db, err := repository.CreateDbConnection(config.Cfg)
	if err != nil {
		config.Logger.Error(err)
	}

	rds, err := repository.CreateRedisConnection(config.Cfg)
	if err != nil {
		config.Logger.Error(err)
	}

	repo := repository.New(db, rds)

	mangaRoute := e.Group("/manga")
	m := mangaservice.New(repo)
	mangacontroller.New(mangaRoute, m)

	linkRoute := e.Group("/link")
	n := linkshortenerservice.New(repo)
	shortenercontroller.New(linkRoute, n)

	bookmarkRoute := e.Group("/bookmark")
	b := bookmarkservice.New(repo)
	bookmarkcontroller.New(bookmarkRoute, b)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome To MangameeApi")
	})

	config.Logger.Info("server start")
	e.Start(":" + config.Cfg.Server.Port)

}
