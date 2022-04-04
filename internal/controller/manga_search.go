package controller

import (
	"mangamee-api/internal/models"
	"mangamee-api/internal/services/scrapper/source_1"
	"mangamee-api/internal/services/scrapper/source_2"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetMangaSearch(c echo.Context) error {

	queryParams := models.QueryParams{
		Source: c.Param("source"),
		Search: strings.Replace(c.QueryParam("title"), " ", "%20", -1),
	}

	switch queryParams.Source {
	case "1":

		mangaData, err := source_1.MangaSearch(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, mangaData)

	case "2":

		mangaData, err := source_2.MangaSearch(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, mangaData)
	}

	return c.JSON(http.StatusBadRequest, "bad request")
}
