package linkcontroller

import (
	"mangamee-api/entity"
	"mangamee-api/exception"
	"mangamee-api/helper"
	"mangamee-api/logger"
	"mangamee-api/response"
	linkservice "mangamee-api/service/link"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type linkControllerImpl struct {
	linkService linkservice.LinkService
}

type LinkHandler interface {
	Mount(group *gin.RouterGroup)
}

func NewLinkController(service linkservice.LinkService) LinkHandler {
	return &linkControllerImpl{
		linkService: service,
	}
}

func (d *linkControllerImpl) Mount(e *gin.RouterGroup) {
	e.POST("", d.InsertLink)
	e.GET("/:id", d.GetLink)
}

func (h *linkControllerImpl) InsertLink(c *gin.Context) {

	ctx := c.Request.Context()

	var requestUrl entity.RequestLink

	if err := c.BindJSON(&requestUrl); err != nil {
		logger.Info("InsertBookmark error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		err = exception.NewBadRequest("bad request")
		response.ErrorRespone(c, err)
		return
	}

	result, err := h.linkService.InsertLink(ctx, requestUrl.Url)
	if err != nil {
		logger.Info("InsertLink error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		response.ErrorRespone(c, err)
		return
	}

	response.SuccessRespone(c, result.Key)
}

func (h *linkControllerImpl) GetLink(c *gin.Context) {

	ctx := c.Request.Context()

	id := c.Param("id")
	result, err := h.linkService.GetLink(ctx, id)
	if err != nil {
		logger.Info("GetLink error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		response.ErrorRespone(c, err)
		return
	}
	response.SuccessRespone(c, result.LongUrl)

}
