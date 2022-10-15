package shortenercontroller

import (
	"context"
	"net/http"

	"mangamee-api/src/config"
	"mangamee-api/src/respone"
	shortenerservice "mangamee-api/src/service/shortener"

	"github.com/labstack/echo/v4"
)

type shortenerControllerImpl struct {
	shortenerService shortenerservice.ShortenerService
}

type ShortenerHandler interface {
	Mount(group *echo.Group)
}

var (
	ctx = context.Background()
)

func NewShortenerController(service shortenerservice.ShortenerService) ShortenerHandler {
	return &shortenerControllerImpl{
		shortenerService: service,
	}
}

func (d *shortenerControllerImpl) Mount(e *echo.Group) {
	e.POST("/:id", d.InsertLink)
	e.GET("/:id", d.GetLink)
}

func (h *shortenerControllerImpl) InsertLink(c echo.Context) error {

	longUrl := c.Param("id")
	result, err := h.shortenerService.InsertLink(ctx, longUrl)
	if err != nil {
		config.Logger.Error(err)
		return respone.JsonError(c, err)
	}

	return respone.JsonSuccess(c, http.StatusCreated, result.Key)
}

func (h *shortenerControllerImpl) GetLink(c echo.Context) error {

	id := c.Param("id")
	result, err := h.shortenerService.GetLink(ctx, id)
	if err != nil {
		config.Logger.Error(err)
		return respone.JsonError(c, err)
	}
	return respone.JsonSuccess(c, http.StatusOK, result.LongUrl)
}
