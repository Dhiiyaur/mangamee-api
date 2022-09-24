package shortenercontroller

import (
	"mangamee-api/internal/config"
	"mangamee-api/internal/respone"
	shortenerservice "mangamee-api/internal/service/shortener"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LinkShortenerHandler struct {
	LinkService *shortenerservice.Service
}

func New(e *echo.Group, s *shortenerservice.Service) {
	handler := &LinkShortenerHandler{
		LinkService: s,
	}

	e.POST("/:id", handler.InsertLink)
	e.GET("/:id", handler.GetLink)
}

func (h *LinkShortenerHandler) InsertLink(c echo.Context) error {
	longUrl := c.Param("id")

	r, err := h.LinkService.GenerateLink(longUrl)
	if err != nil {
		config.Logger.Error(err)
		return respone.JsonError(c, http.StatusBadRequest, err.Error())
	}

	return respone.JsonSuccess(c, http.StatusCreated, r)
}

func (h *LinkShortenerHandler) GetLink(c echo.Context) error {
	id := c.Param("id")

	r, err := h.LinkService.GetLongUrl(id)
	if err != nil {
		config.Logger.Error(err)
		return respone.JsonError(c, http.StatusBadRequest, err.Error())
	}

	return respone.JsonSuccess(c, http.StatusOK, r)
}
