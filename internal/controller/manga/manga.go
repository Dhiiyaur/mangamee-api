package mangacontroller

import (
	"mangamee-api/internal/config"
	"mangamee-api/internal/entity"
	"mangamee-api/internal/respone"
	mangaservice "mangamee-api/internal/service/manga"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type MangaHandler struct {
	MangaService *mangaservice.Service
}

func New(e *echo.Group, s *mangaservice.Service) {
	handler := &MangaHandler{
		MangaService: s,
	}

	e.GET("/index/:source/:page", handler.GetIndex)
	e.GET("/search/:source", handler.GetSearch)
	e.GET("/detail/:source/:id", handler.GetDetail)
	e.GET("/read/:source/:id/:chapter_id", handler.GetImage)
	e.GET("/chapter/:source/:id", handler.GetChapter)
	e.GET("/source", handler.GetSource)
	e.GET("/meta/:source/:id", handler.GetMetaTag)
}

func (h *MangaHandler) GetIndex(c echo.Context) error {
	params := entity.MangaParams{
		Source:     c.Param("source"),
		PageNumber: c.Param("page"),
		Path:       "index",
	}

	r, err := h.MangaService.GetIndex(params)
	if err != nil {
		config.Logger.Error(err)
		respone.JsonError(c, http.StatusBadRequest, err.Error())
	}

	return respone.JsonSuccess(c, http.StatusOK, r)
}

func (h *MangaHandler) GetSearch(c echo.Context) error {
	params := entity.MangaParams{
		Source: c.Param("source"),
		Search: strings.Replace(c.QueryParam("title"), " ", "%20", -1),
		Path:   "search",
	}

	r, err := h.MangaService.GetSearch(params)
	if err != nil {
		config.Logger.Error(err)
		respone.JsonError(c, http.StatusBadRequest, err.Error())
	}

	return respone.JsonSuccess(c, http.StatusOK, r)
}

func (h *MangaHandler) GetDetail(c echo.Context) error {
	params := entity.MangaParams{
		Source:  c.Param("source"),
		MangaId: c.Param("id"),
		Path:    "detail",
	}

	r, err := h.MangaService.GetDetail(params)
	if err != nil {
		config.Logger.Error(err)
		respone.JsonError(c, http.StatusBadRequest, err.Error())
	}

	return respone.JsonSuccess(c, http.StatusOK, r)
}

func (h *MangaHandler) GetImage(c echo.Context) error {
	params := entity.MangaParams{
		Source:    c.Param("source"),
		MangaId:   c.Param("id"),
		ChapterId: c.Param("chapter_id"),
		Path:      "image",
	}

	r, err := h.MangaService.GetImage(params)
	if err != nil {
		config.Logger.Error(err)
		respone.JsonError(c, http.StatusBadRequest, err.Error())
	}

	return respone.JsonSuccess(c, http.StatusOK, r)
}

func (h *MangaHandler) GetChapter(c echo.Context) error {
	params := entity.MangaParams{
		Source:  c.Param("source"),
		MangaId: c.Param("id"),
		Path:    "chapter",
	}

	r, err := h.MangaService.GetChapter(params)
	if err != nil {
		config.Logger.Error(err)
		respone.JsonError(c, http.StatusBadRequest, err.Error())
	}

	return respone.JsonSuccess(c, http.StatusOK, r)
}

func (h *MangaHandler) GetMetaTag(c echo.Context) error {
	params := entity.MangaParams{
		Source:  c.Param("source"),
		MangaId: c.Param("id"),
		Path:    "meta",
	}

	r, err := h.MangaService.GetMeta(params)
	if err != nil {
		config.Logger.Error(err)
		respone.JsonError(c, http.StatusBadRequest, err.Error())
	}

	return respone.JsonSuccess(c, http.StatusOK, r)
}

func (h *MangaHandler) GetSource(c echo.Context) error {

	MangaSource := []entity.MangaSource{
		{
			Name: "Mangaread",
			Id:   1,
		},
		{
			Name: "Mangatown",
			Id:   2,
		},
		{
			Name: "Mangabat",
			Id:   3,
		},
		{
			Name: "Maidmy",
			Id:   4,
		},
	}

	return respone.JsonSuccess(c, http.StatusOK, MangaSource)
}
