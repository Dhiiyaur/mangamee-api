package bookmarkcontroller

import (
	"encoding/json"
	"mangamee-api/internal/respone"
	bookmarkservice "mangamee-api/internal/service/bookmark"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookmarkHandler struct {
	BookmarkService *bookmarkservice.Service
}

func New(e *echo.Group, s *bookmarkservice.Service) {
	handler := &BookmarkHandler{
		BookmarkService: s,
	}

	e.POST("/", handler.GetCode)
	e.GET("/:id", handler.GetBookmark)
}

func (h *BookmarkHandler) GetCode(c echo.Context) error {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return respone.JsonError(c, http.StatusBadRequest, err.Error())
	}

	r, err := h.BookmarkService.GenerateCode(json_map["data"])

	if err != nil {
		return respone.JsonError(c, http.StatusBadRequest, err.Error())
	}

	return respone.JsonSuccess(c, http.StatusCreated, r)
}

func (h *BookmarkHandler) GetBookmark(c echo.Context) error {

	code := c.Param("id")
	bookmarkData, err := h.BookmarkService.GetBookmark(code)
	if err != nil {
		return respone.JsonError(c, http.StatusBadRequest, err.Error())
	}

	return respone.JsonSuccess(c, http.StatusOK, bookmarkData)

}
