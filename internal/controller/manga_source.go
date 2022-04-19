package controller

import (
	"mangamee-api/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetMangaSource(c echo.Context) error {

	var MangaSource []models.MangaSource

	var source_1 = models.MangaSource{
		Name: "Mangaread",
		Id:   1,
	}

	var source_2 = models.MangaSource{
		Name: "Mangatown",
		Id:   2,
	}

	var source_3 = models.MangaSource{
		Name: "Maidmy",
		Id:   3,
	}

	MangaSource = append(MangaSource, source_1)
	MangaSource = append(MangaSource, source_2)
	MangaSource = append(MangaSource, source_3)

	return c.JSON(http.StatusOK, MangaSource)
}
