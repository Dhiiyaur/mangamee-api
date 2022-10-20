package response

import (
	"mangamee-api/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessRespone(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    data,
	})
}

func ErrorRespone(c *gin.Context, err error) {
	c.JSON(exception.Status(err), gin.H{
		"message": err.Error(),
	})
}
