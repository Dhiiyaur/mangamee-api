package app

import (
	"mangamee-api/internal/config"
	mangacontroller "mangamee-api/internal/controller/manga"
	"mangamee-api/internal/repository"
	mangaservice "mangamee-api/internal/service/manga"

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

	m := mangaservice.New(repo)
	mangacontroller.New(e, m)

	config.Logger.Info("server start")
	e.Start(":" + config.Cfg.Server.Port)

}
