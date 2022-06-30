package controller

import (
	"mangamee-api/internal/db"
	"mangamee-api/internal/models"
	"mangamee-api/internal/services/scrapper/source_1"
	"mangamee-api/internal/services/scrapper/source_2"
	"mangamee-api/internal/services/scrapper/source_3"
	"mangamee-api/internal/services/scrapper/source_4"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetMangaSearch(c echo.Context) error {

	queryParams := models.QueryParams{
		Source: c.Param("source"),
		Search: strings.Replace(c.QueryParam("title"), " ", "%20", -1),
	}

	MangaData := models.ReturnData{}
	cache, err := db.CacheChecking("search", queryParams)
	if err == nil {
		return c.JSON(http.StatusOK, cache.Datas)
	}

	switch queryParams.Source {
	case "1":

		db.InsertDataUserLog("search", 1, queryParams.Search, "-")
		MangaData.Datas, err = source_1.MangaSearch(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		db.SetCache("search", queryParams, MangaData)
		return c.JSON(http.StatusOK, MangaData.Datas)

	case "2":

		db.InsertDataUserLog("search", 2, queryParams.Search, "-")
		MangaData.Datas, err = source_2.MangaSearch(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		db.SetCache("search", queryParams, MangaData)
		return c.JSON(http.StatusOK, MangaData.Datas)

	case "3":

		db.InsertDataUserLog("search", 3, queryParams.Search, "-")
		MangaData.Datas, err = source_3.MangaSearch(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		db.SetCache("search", queryParams, MangaData)
		return c.JSON(http.StatusOK, MangaData.Datas)

	case "4":

		db.InsertDataUserLog("search", 4, queryParams.Search, "-")
		MangaData.Datas, err = source_4.MangaSearch(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		db.SetCache("search", queryParams, MangaData)
		return c.JSON(http.StatusOK, MangaData.Datas)

	}
	return c.JSON(http.StatusBadRequest, "bad request")
}
