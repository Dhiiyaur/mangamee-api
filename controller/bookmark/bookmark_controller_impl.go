package bookmarkcontroller

import (
	"encoding/json"
	"mangamee-api/exception"
	"mangamee-api/helper"
	"mangamee-api/logger"
	"mangamee-api/response"

	bookmarkservice "mangamee-api/service/bookmark"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type bookmarkControllerImpl struct {
	bookmarkService bookmarkservice.BookmarkService
}

type BookmarkHandler interface {
	Mount(group *gin.RouterGroup)
}

func NewBookmarkController(service bookmarkservice.BookmarkService) BookmarkHandler {
	return &bookmarkControllerImpl{
		bookmarkService: service,
	}
}

func (d *bookmarkControllerImpl) Mount(e *gin.RouterGroup) {

	e.POST("/", d.InsertBookmark)
	e.GET("/:id", d.GetBookmark)
}

func (h *bookmarkControllerImpl) InsertBookmark(c *gin.Context) {

	ctx := c.Request.Context()

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)
	if err != nil {
		if err != nil {
			logger.Info("InsertBookmark error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
			err = exception.NewBadRequest("bad request")
			response.ErrorRespone(c, err)
			return
		}
	}

	result, err := h.bookmarkService.InsertBookmark(ctx, json_map["data"])
	if err != nil {
		logger.Info("InsertBookmark error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		response.ErrorRespone(c, err)
		return
	}

	response.SuccessRespone(c, result)
}

func (h *bookmarkControllerImpl) GetBookmark(c *gin.Context) {

	ctx := c.Request.Context()

	bookmarkId := c.Param("id")
	result, err := h.bookmarkService.GetBookmark(ctx, bookmarkId)
	if err != nil {
		logger.Info("GetBookmark error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		response.ErrorRespone(c, err)
		return
	}

	response.SuccessRespone(c, result)
}
