package app

import (
	"os"

	"github.com/labstack/echo/v4"
)

func Start() {

	e := echo.New()
	MiddlewareSetup(e)
	RouterApp(e)

	e.Start(":" + os.Getenv("PORT"))
}
