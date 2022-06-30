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

func GetMangaDetail(c echo.Context) error {

	queryParams := models.QueryParams{
		Source: c.Param("source"),
		Id:     c.Param("id"),
	}

	MangaData := models.ReturnData{}

	cache, err := db.CacheChecking("detail", queryParams)
	if err == nil {
		return c.JSON(http.StatusOK, cache.Data)
	}

	switch queryParams.Source {
	case "1":

		db.InsertDataUserLog("detail", 1, queryParams.Id, "-")
		MangaData.Data, err = source_1.MangaDetail(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		db.SetCache("detail", queryParams, MangaData)
		return c.JSON(http.StatusOK, MangaData.Data)

	case "2":

		db.InsertDataUserLog("detail", 2, queryParams.Id, "-")
		MangaData.Data, err = source_2.MangaDetail(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		db.SetCache("detail", queryParams, MangaData)
		return c.JSON(http.StatusOK, MangaData.Data)

	case "3":

		db.InsertDataUserLog("detail", 3, queryParams.Id, "-")
		MangaData.Data, err = source_3.MangaDetail(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		db.SetCache("detail", queryParams, MangaData)
		return c.JSON(http.StatusOK, MangaData.Data)

	case "4":

		db.InsertDataUserLog("detail", 4, queryParams.Id, "-")
		MangaData.Data, err = source_4.MangaDetail(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		db.SetCache("detail", queryParams, MangaData)
		return c.JSON(http.StatusOK, MangaData.Data)

	}

	return c.JSON(http.StatusBadRequest, "bad request")
}

func GetMangaMetaTag(c echo.Context) error {

	queryParams := models.QueryParams{
		Source: c.Param("source"),
		Id:     c.Param("id"),
	}

	MangaData := models.ReturnData{}
	var err error

	switch queryParams.Source {
	case "1":

		MangaData.Data, err = source_1.MangaMetaTag(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, MangaData.Data)

	case "2":

		MangaData.Data, err = source_2.MangaMetaTag(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, MangaData.Data)

	case "3":

		MangaData.Data, err = source_3.MangaMetaTag(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, MangaData.Data)

	case "4":

		MangaData.Data, err = source_4.MangaMetaTag(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, MangaData.Data)

	}

	return c.JSON(http.StatusBadRequest, "bad request")

}
