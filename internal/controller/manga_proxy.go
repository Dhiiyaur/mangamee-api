package controller

import (
	"mangamee-api/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func getSourceHeader(queryParams models.QueryParams) string {

	var headerSet string

	switch queryParams.Source {
	case "4":
		headerSet = "https://m.mangabat.com/"
	}

	return headerSet

}

func GetMangaProxy(c echo.Context) error {

	queryParams := models.QueryParams{
		Source:     c.Param("source"),
		ImageProxy: c.QueryParam("id"),
	}

	req, err := http.NewRequest("GET", queryParams.ImageProxy, nil)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	req.Header.Set("Referer", getSourceHeader(queryParams))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.Stream(http.StatusOK, "image", resp.Body)

}
