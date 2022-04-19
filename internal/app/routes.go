package app

import (
	"net/http"

	"mangamee-api/internal/controller"

	"github.com/labstack/echo/v4"
)

func RouterApp(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome To MangameeApi")
	})

	//  Manga

	m := e.Group("/manga")
	m.GET("/index/:source/:page", controller.GetMangaIndex)
	m.GET("/detail/:source/:id", controller.GetMangaDetail)
	m.GET("/read/:source/:id/:chapter_id", controller.GetMangaImage)
	m.GET("/read-chapter/:source/:id", controller.GetMangaChapther)
	m.GET("/search/:source", controller.GetMangaSearch)
	m.GET("/meta/:source/:id", controller.GetMangaMetaTag)
	m.GET("/source", controller.GetMangaSource)

	// auth

	// e.GET("/editorpick", controller.EditorPick)
	// e.GET("/browse", controller.Browse)
	// e.GET("/search", controller.Search)

	// e.GET("analytics/all", controller.GetCountApiHit)
}
