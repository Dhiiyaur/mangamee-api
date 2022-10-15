package mangacontroller

import "github.com/labstack/echo/v4"

type MangaController interface {
	GetIndex(c echo.Context) error
	GetSearch(c echo.Context) error
	GetDetail(c echo.Context) error
	GetImage(c echo.Context) error
	GetChapter(c echo.Context) error
	GetSource(c echo.Context) error
	GetMetaTag(c echo.Context) error
	GetMangaProxy(c echo.Context) error
}
