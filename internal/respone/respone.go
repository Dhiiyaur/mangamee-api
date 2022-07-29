package respone

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	SuccessResponse struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	ErrorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

func JsonSuccess(c echo.Context, code int, data interface{}) error {
	return c.JSON(http.StatusOK, SuccessResponse{
		Code:    code,
		Message: "success",
		Data:    data,
	})
}

func JsonError(c echo.Context, code int, message string) error {
	return c.JSON(code, ErrorResponse{
		Code:    code,
		Message: message,
	})
}
