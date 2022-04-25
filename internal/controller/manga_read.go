package controller

import (
	"net/http"

	"mangamee-api/internal/db"
	"mangamee-api/internal/models"
	"mangamee-api/internal/services/scrapper/source_1"
	"mangamee-api/internal/services/scrapper/source_2"
	"mangamee-api/internal/services/scrapper/source_3"

	"github.com/labstack/echo/v4"
)

func GetMangaImage(c echo.Context) error {

	queryParams := models.QueryParams{
		Source:    c.Param("source"),
		Id:        c.Param("id"),
		ChapterId: c.Param("chapter_id"),
	}

	switch queryParams.Source {
	case "1":

		db.InsertDataUserLog("read", "1", queryParams.Id, queryParams.ChapterId, c.RealIP())
		mangaData, err := source_1.MangaImage(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, mangaData)

	case "2":

		db.InsertDataUserLog("read", "2", queryParams.Id, queryParams.ChapterId, c.RealIP())
		mangaData, err := source_2.MangaImage(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, mangaData)

	case "3":

		db.InsertDataUserLog("read", "3", queryParams.Id, queryParams.ChapterId, c.RealIP())
		mangaData, err := source_3.MangaImage(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, mangaData)
	}

	return c.JSON(http.StatusBadRequest, "bad request")

}

func GetMangaChapther(c echo.Context) error {

	queryParams := models.QueryParams{
		Source: c.Param("source"),
		Id:     c.Param("id"),
	}

	switch queryParams.Source {
	case "1":

		mangaData, err := source_1.MangaChapter(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, mangaData)

	case "2":

		mangaData, err := source_2.MangaChapter(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, mangaData)

	case "3":

		mangaData, err := source_3.MangaChapter(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, mangaData)

	}

	return c.JSON(http.StatusBadRequest, "bad request")
}
