package app

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func Start() {

	e := echo.New()
	MiddlewareSetup(e)
	RouterApp(e)
	e.Start(":" + viper.GetString("PORT"))
}
