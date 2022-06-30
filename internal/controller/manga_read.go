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

func GetMangaImage(c echo.Context) error {

	queryParams := models.QueryParams{
		Source:    c.Param("source"),
		Id:        c.Param("id"),
		ChapterId: c.Param("chapter_id"),
	}

	MangaData := models.ReturnData{}

	cache, err := db.CacheChecking("read", queryParams)
	if err == nil {
		return c.JSON(http.StatusOK, cache.Data)
	}

	switch queryParams.Source {
	case "1":

		db.InsertDataUserLog("read", 1, queryParams.Id, queryParams.ChapterId)
		MangaData.Data, err = source_1.MangaImage(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		db.SetCache("read", queryParams, MangaData)
		return c.JSON(http.StatusOK, MangaData.Data)

	case "2":

		db.InsertDataUserLog("read", 2, queryParams.Id, queryParams.ChapterId)
		MangaData.Data, err = source_2.MangaImage(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		db.SetCache("read", queryParams, MangaData)
		return c.JSON(http.StatusOK, MangaData.Data)

	case "3":

		db.InsertDataUserLog("read", 3, queryParams.Id, queryParams.ChapterId)
		MangaData.Data, err = source_3.MangaImage(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		db.SetCache("read", queryParams, MangaData)
		return c.JSON(http.StatusOK, MangaData.Data)

	case "4":

		db.InsertDataUserLog("read", 4, queryParams.Id, queryParams.ChapterId)
		MangaData.Data, err = source_4.MangaImage(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		db.SetCache("read", queryParams, MangaData)
		return c.JSON(http.StatusOK, MangaData.Data)
	}

	return c.JSON(http.StatusBadRequest, "bad request")

}

func GetMangaChapther(c echo.Context) error {

	queryParams := models.QueryParams{
		Source: c.Param("source"),
		Id:     c.Param("id"),
	}

	MangaData := models.ReturnData{}

	cache, err := db.CacheChecking("chapter", queryParams)
	if err == nil {
		return c.JSON(http.StatusOK, cache.Data)
	}

	switch queryParams.Source {
	case "1":

		MangaData.Data, err = source_1.MangaChapter(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		db.SetCache("chapter", queryParams, MangaData)
		return c.JSON(http.StatusOK, MangaData.Data)

	case "2":

		MangaData.Data, err = source_2.MangaChapter(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		db.SetCache("chapter", queryParams, MangaData)
		return c.JSON(http.StatusOK, MangaData.Data)

	case "3":

		MangaData.Data, err = source_3.MangaChapter(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		db.SetCache("chapter", queryParams, MangaData)
		return c.JSON(http.StatusOK, MangaData.Data)

	case "4":

		MangaData.Data, err = source_4.MangaChapter(queryParams)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		db.SetCache("chapter", queryParams, MangaData)
		return c.JSON(http.StatusOK, MangaData.Data)
	}

	return c.JSON(http.StatusBadRequest, "bad request")
}
