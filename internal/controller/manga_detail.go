package controller

import (
	"net/http"

	"mangamee-api/internal/models"
	"mangamee-api/internal/services/scrapper/source_1"
	"mangamee-api/internal/services/scrapper/source_2"
	"mangamee-api/internal/services/scrapper/source_3"

	"github.com/labstack/echo/v4"
)

func GetMangaDetail(c echo.Context) error {

	queryParams := models.QueryParams{
		Source: c.Param("source"),
		Id:     c.Param("id"),
	}

	switch queryParams.Source {
	case "1":

		mangaData, err := source_1.MangaDetail(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, mangaData)

	case "2":

		mangaData, err := source_2.MangaDetail(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, mangaData)

	case "3":

		mangaData, err := source_3.MangaDetail(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, mangaData)
	}

	return c.JSON(http.StatusBadRequest, "bad request")
}

func GetMangaMetaTag(c echo.Context) error {
	queryParams := models.QueryParams{
		Source: c.Param("source"),
		Id:     c.Param("id"),
	}

	switch queryParams.Source {
	case "1":

		mangaData, err := source_1.MangaMetaTag(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, mangaData)

	case "2":

		mangaData, err := source_2.MangaMetaTag(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, mangaData)

	case "3":

		mangaData, err := source_3.MangaMetaTag(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, mangaData)
	}

	return c.JSON(http.StatusBadRequest, "bad request")

}
