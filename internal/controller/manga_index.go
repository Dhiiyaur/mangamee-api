package controller

import (
	"net/http"

	"mangamee-api/internal/db"
	"mangamee-api/internal/models"
	"mangamee-api/internal/services/scrapper/source_1"
	"mangamee-api/internal/services/scrapper/source_2"
	"mangamee-api/internal/services/scrapper/source_3"
	"mangamee-api/internal/services/scrapper/source_4"

	"github.com/labstack/echo/v4"
)

func GetMangaIndex(c echo.Context) error {

	queryParams := models.QueryParams{
		Source: c.Param("source"),
		Page:   c.Param("page"),
	}

	switch queryParams.Source {
	case "1":

		db.InsertDataUserLog("index", 1, "-", "-")
		mangaData, err := source_1.MangaIndex(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, mangaData)

	case "2":

		db.InsertDataUserLog("index", 2, "-", "-")
		mangaData, err := source_2.MangaIndex(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, mangaData)

	case "3":

		db.InsertDataUserLog("index", 3, "-", "-")
		mangaData, err := source_3.MangaIndex(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, mangaData)

	case "4":

		db.InsertDataUserLog("index", 4, "-", "-")
		mangaData, err := source_4.MangaIndex(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, mangaData)

	}

	return c.JSON(http.StatusBadRequest, "bad request")

}
