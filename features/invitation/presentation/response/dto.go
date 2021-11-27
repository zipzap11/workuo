package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string      `json: "message"`
	Data    interface{} `json: "data"`
}

func NewSuccessResponse(e echo.Context, data interface{}) error {
	return e.JSON(http.StatusOK, Response{
		Message: "Success",
		Data:    data,
	})
}

func NewErrorResponse(e echo.Context, code int, msg string) error {
	return e.JSON(code, Response{
		Message: msg,
	})
}
