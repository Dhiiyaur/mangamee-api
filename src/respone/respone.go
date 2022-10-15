package respone

import (
	"mangamee-api/src/exception"

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
	return c.JSON(code, SuccessResponse{
		Code:    code,
		Message: "success",
		Data:    data,
	})
}

func JsonError(c echo.Context, err error) error {
	switch err := err.(type) {
	case exception.ErrorWithCode:
		res := ErrorResponse{
			Code:    err.HttpResponeCode,
			Message: err.Msg,
		}
		return c.JSON(err.HttpResponeCode, res)
	default:
		res := ErrorResponse{
			Code:    500,
			Message: "internal server error",
		}
		return c.JSON(500, res)
	}

}
