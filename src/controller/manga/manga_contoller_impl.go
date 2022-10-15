package mangacontroller

import (
	"context"
	"mangamee-api/src/config"
	"mangamee-api/src/entity"
	"mangamee-api/src/exception"
	"mangamee-api/src/respone"
	mangaservice "mangamee-api/src/service/manga"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type mangaControllerImpl struct {
	mangaService mangaservice.MangaService
}

type MangaHandler interface {
	Mount(group *echo.Group)
}

var (
	ctx = context.Background()
)

func NewMangaController(service mangaservice.MangaService) MangaHandler {
	return &mangaControllerImpl{
		mangaService: service,
	}
}

func (d *mangaControllerImpl) Mount(e *echo.Group) {
	e.GET("/index/:source/:page", d.GetIndex)
	e.GET("/search/:source", d.GetSearch)
	e.GET("/detail/:source/:id", d.GetDetail)
	e.GET("/read/:source/:id/:chapter_id", d.GetImage)
	e.GET("/chapter/:source/:id", d.GetChapter)
	e.GET("/source", d.GetSource)
	e.GET("/meta/:source/:id", d.GetMetaTag)
	e.GET("/proxy", d.GetMangaProxy)
}

func (h *mangaControllerImpl) GetIndex(c echo.Context) error {
	params := entity.MangaParams{
		Source:     c.Param("source"),
		PageNumber: c.Param("page"),
		Path:       "index",
	}

	result, err := h.mangaService.GetIndex(ctx, params)
	if err != nil {
		config.Logger.Error(err)
		return respone.JsonError(c, err)
	}

	return respone.JsonSuccess(c, http.StatusOK, result)
}

func (h *mangaControllerImpl) GetSearch(c echo.Context) error {
	params := entity.MangaParams{
		Source: c.Param("source"),
		Search: strings.Replace(c.QueryParam("title"), " ", "%20", -1),
		Path:   "search",
	}

	result, err := h.mangaService.GetSearch(ctx, params)
	if err != nil {
		config.Logger.Error(err)
		return respone.JsonError(c, err)
	}

	return respone.JsonSuccess(c, http.StatusOK, result)

}

func (h *mangaControllerImpl) GetDetail(c echo.Context) error {
	params := entity.MangaParams{
		Source:  c.Param("source"),
		MangaId: c.Param("id"),
		Path:    "detail",
	}

	result, err := h.mangaService.GetDetail(ctx, params)
	if err != nil {
		config.Logger.Error(err)
		return respone.JsonError(c, err)
	}
	return respone.JsonSuccess(c, http.StatusOK, result)
}

func (h *mangaControllerImpl) GetImage(c echo.Context) error {
	params := entity.MangaParams{
		Source:    c.Param("source"),
		MangaId:   c.Param("id"),
		ChapterId: c.Param("chapter_id"),
		Path:      "image",
	}

	result, err := h.mangaService.GetImage(ctx, params)
	if err != nil {
		config.Logger.Error(err)
		return respone.JsonError(c, err)
	}

	return respone.JsonSuccess(c, http.StatusOK, result)
}

func (h *mangaControllerImpl) GetChapter(c echo.Context) error {
	params := entity.MangaParams{
		Source:  c.Param("source"),
		MangaId: c.Param("id"),
		Path:    "chapter",
	}

	result, err := h.mangaService.GetChapter(ctx, params)
	if err != nil {
		config.Logger.Error(err)
		return respone.JsonError(c, err)
	}

	return respone.JsonSuccess(c, http.StatusOK, result)
}

func (h *mangaControllerImpl) GetMetaTag(c echo.Context) error {
	params := entity.MangaParams{
		Source:  c.Param("source"),
		MangaId: c.Param("id"),
		Path:    "meta",
	}

	result, err := h.mangaService.GetMeta(ctx, params)
	if err != nil {
		config.Logger.Error(err)
		return respone.JsonError(c, err)
	}

	return respone.JsonSuccess(c, http.StatusOK, result)
}

func (h *mangaControllerImpl) GetSource(c echo.Context) error {

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

func (h *mangaControllerImpl) GetMangaProxy(c echo.Context) error {
	params := entity.MangaParams{
		ImageProxy: c.QueryParam("id"),
	}
	req, err := http.NewRequest("GET", params.ImageProxy, nil)
	if err != nil {
		config.Logger.Error(err)
		err = exception.NewErrorMsg(
			exception.CodeBadRequest, err,
		)
		return respone.JsonError(c, err)
	}

	req.Header.Set("Referer", "https://m.mangabat.com/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		config.Logger.Error(err)
		err = exception.NewErrorMsg(
			exception.CodeBadRequest, err,
		)
		return respone.JsonError(c, err)
	}

	return c.Stream(http.StatusOK, "image", resp.Body)
}
