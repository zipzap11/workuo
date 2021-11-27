package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string
	Data    interface{}
}

func NewErrorResponse(e echo.Context, msg string, code int) error {
	return e.JSON(code, Response{
		Message: msg,
	})
}

func NewSuccessResponse(e echo.Context, msg string, data interface{}) error {
	return e.JSON(http.StatusOK, Response{
		Message: msg,
		Data:    data,
	})
}
