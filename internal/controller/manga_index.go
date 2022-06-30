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

	MangaData := models.ReturnData{}

	cache, err := db.CacheChecking("index", queryParams)
	if err == nil {
		return c.JSON(http.StatusOK, cache.Datas)
	}

	switch queryParams.Source {
	case "1":

		db.InsertDataUserLog("index", 1, "-", "-")
		MangaData.Datas, err = source_1.MangaIndex(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		db.SetCache("index", queryParams, MangaData)
		return c.JSON(http.StatusOK, MangaData.Datas)

	case "2":

		db.InsertDataUserLog("index", 2, "-", "-")
		MangaData.Datas, err = source_2.MangaIndex(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		db.SetCache("index", queryParams, MangaData)
		return c.JSON(http.StatusOK, MangaData.Datas)

	case "3":

		db.InsertDataUserLog("index", 3, "-", "-")
		MangaData.Datas, err = source_3.MangaIndex(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		db.SetCache("index", queryParams, MangaData)
		return c.JSON(http.StatusOK, MangaData.Datas)

	case "4":

		db.InsertDataUserLog("index", 4, "-", "-")
		MangaData.Datas, err = source_4.MangaIndex(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		db.SetCache("index", queryParams, MangaData)
		return c.JSON(http.StatusOK, MangaData.Datas)

	}

	return c.JSON(http.StatusBadRequest, "bad request")

}
