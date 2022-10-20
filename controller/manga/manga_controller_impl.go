package mangacontroller

import (
	"io/ioutil"
	"mangamee-api/entity"
	"mangamee-api/exception"
	"mangamee-api/helper"
	"mangamee-api/logger"
	"mangamee-api/response"
	mangaservice "mangamee-api/service/manga"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type mangaControllerImpl struct {
	mangaService mangaservice.MangaService
}

type MangaHandler interface {
	Mount(group *gin.RouterGroup)
}

func NewMangaController(service mangaservice.MangaService) MangaHandler {
	return &mangaControllerImpl{
		mangaService: service,
	}
}

func (d *mangaControllerImpl) Mount(e *gin.RouterGroup) {
	e.GET("/index/:source/:page", d.GetIndex)
	e.GET("/search/:source", d.GetSearch)
	e.GET("/detail/:source/:id", d.GetDetail)
	e.GET("/read/:source/:id/:chapter_id", d.GetImage)
	e.GET("/chapter/:source/:id", d.GetChapter)
	e.GET("/source", d.GetSource)
	e.GET("/meta/:source/:id", d.GetMetaTag)
	e.GET("/proxy", d.GetMangaProxy)
}

func (h *mangaControllerImpl) GetIndex(c *gin.Context) {

	ctx := c.Request.Context()
	params := entity.RequestParams{
		Source:     c.Param("source"),
		PageNumber: c.Param("page"),
		Path:       "index",
	}
	result, err := h.mangaService.GetIndex(ctx, params)
	if err != nil {
		logger.Info("GetIndex error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		response.ErrorRespone(c, err)
		return
	}
	response.SuccessRespone(c, result)
}

func (h *mangaControllerImpl) GetSearch(c *gin.Context) {
	ctx := c.Request.Context()
	params := entity.RequestParams{
		Source: c.Param("source"),
		Search: strings.Replace(c.Query("title"), " ", "%20", -1),
		Path:   "search",
	}

	result, err := h.mangaService.GetSearch(ctx, params)
	if err != nil {
		logger.Info("GetSearch error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		response.ErrorRespone(c, err)
		return
	}

	response.SuccessRespone(c, result)

}

func (h *mangaControllerImpl) GetDetail(c *gin.Context) {
	ctx := c.Request.Context()
	params := entity.RequestParams{
		Source:  c.Param("source"),
		MangaId: c.Param("id"),
		Path:    "detail",
	}

	result, err := h.mangaService.GetDetail(ctx, params)
	if err != nil {
		logger.Info("GetDetail error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		response.ErrorRespone(c, err)
		return
	}

	response.SuccessRespone(c, result)
}

func (h *mangaControllerImpl) GetImage(c *gin.Context) {
	ctx := c.Request.Context()
	params := entity.RequestParams{
		Source:    c.Param("source"),
		MangaId:   c.Param("id"),
		ChapterId: c.Param("chapter_id"),
		Path:      "image",
	}

	result, err := h.mangaService.GetImage(ctx, params)
	if err != nil {
		logger.Info("GetDetail error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		response.ErrorRespone(c, err)
		return
	}

	response.SuccessRespone(c, result)
}

func (h *mangaControllerImpl) GetChapter(c *gin.Context) {

	ctx := c.Request.Context()
	params := entity.RequestParams{
		Source:  c.Param("source"),
		MangaId: c.Param("id"),
		Path:    "chapter",
	}

	result, err := h.mangaService.GetChapter(ctx, params)
	if err != nil {
		logger.Info("GetChapter error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		response.ErrorRespone(c, err)
		return
	}

	response.SuccessRespone(c, result)
}

func (h *mangaControllerImpl) GetMetaTag(c *gin.Context) {
	ctx := c.Request.Context()

	params := entity.RequestParams{
		Source:  c.Param("source"),
		MangaId: c.Param("id"),
		Path:    "meta",
	}

	result, err := h.mangaService.GetMeta(ctx, params)
	if err != nil {
		logger.Info("GetMetaTag error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		response.ErrorRespone(c, err)
		return
	}

	response.SuccessRespone(c, result)
}

func (h *mangaControllerImpl) GetSource(c *gin.Context) {

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

	response.SuccessRespone(c, MangaSource)
}

func (h *mangaControllerImpl) GetMangaProxy(c *gin.Context) {
	ctx := c.Request.Context()

	params := entity.RequestParams{
		ImageProxy: c.Query("id"),
	}

	req, err := http.NewRequest("GET", params.ImageProxy, nil)
	if err != nil {
		logger.Info("GetMangaProxy error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		err = exception.NewBadRequest("bad request")
		response.ErrorRespone(c, err)
		return
	}

	req.Header.Set("Referer", "https://m.mangabat.com/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Info("GetMangaProxy error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		err = exception.NewBadRequest("bad request")
		response.ErrorRespone(c, err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Info("GetMangaProxy error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		err = exception.NewBadRequest("bad request")
		response.ErrorRespone(c, err)
		return
	}

	c.Writer.Write(body)
}
