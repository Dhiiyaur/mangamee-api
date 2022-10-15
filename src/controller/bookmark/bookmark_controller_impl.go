package bookmarkcontroller

import (
	"context"
	"encoding/json"
	"net/http"

	"mangamee-api/src/config"
	"mangamee-api/src/exception"
	"mangamee-api/src/respone"
	bookmarkservice "mangamee-api/src/service/bookmark"

	"github.com/labstack/echo/v4"
)

type bookmarkControllerImpl struct {
	bookmarkService bookmarkservice.BookmarkService
}

type BookmarkHandler interface {
	Mount(group *echo.Group)
}

var (
	ctx = context.Background()
)

func NewBookmarkController(service bookmarkservice.BookmarkService) BookmarkHandler {
	return &bookmarkControllerImpl{
		bookmarkService: service,
	}
}

func (d *bookmarkControllerImpl) Mount(e *echo.Group) {

	e.POST("/", d.InsertBookmark)
	e.GET("/:id", d.GetBookmark)
}

func (h *bookmarkControllerImpl) InsertBookmark(c echo.Context) error {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		config.Logger.Error(err)
		err = exception.NewErrorMsg(
			exception.CodeBadRequest, err,
		)
		return respone.JsonError(c, err)
	}

	result, err := h.bookmarkService.InsertBookmark(ctx, json_map["data"])
	if err != nil {
		config.Logger.Error(err)
		return respone.JsonError(c, err)
	}
	return respone.JsonSuccess(c, http.StatusCreated, result)
}

func (h *bookmarkControllerImpl) GetBookmark(c echo.Context) error {

	bookmarkId := c.Param("id")
	result, err := h.bookmarkService.GetBookmark(ctx, bookmarkId)
	if err != nil {
		config.Logger.Error(err)
		return respone.JsonError(c, err)
	}
	return respone.JsonSuccess(c, http.StatusOK, result)
}
