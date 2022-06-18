package controller

import (
	"mangamee-api/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetMangaSource(c echo.Context) error {

	MangaSource := []models.MangaSource{
		{
			Name: "Mangaread",
			Id:   1,
		},
		{
			Name: "Mangatown",
			Id:   2,
		},
		{
			Name: "Maidmy",
			Id:   3,
		},
		{
			Name: "Mangabat",
			Id:   4,
		},
	}

	return c.JSON(http.StatusOK, MangaSource)
}
