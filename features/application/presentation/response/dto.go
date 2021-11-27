package response

import (
	"net/http"
	"workuo/features/application"

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

type ApplicationResponse struct {
	ID     uint
	UserID uint
	JobID  uint
	Job    JobResponse
}

type JobResponse struct {
	ID          uint   `json: "id"`
	Title       string `json: "title"`
	Description string `json: "description`
}

func ToApplicationResponse(data application.ApplicationCore) ApplicationResponse {
	return ApplicationResponse{
		ID:     data.ID,
		UserID: data.UserID,
		JobID:  data.JobID,
		Job:    ToJobResponse(data.Job),
	}
}

func ToJobResponse(data application.JobCore) JobResponse {
	return JobResponse{
		ID:          uint(data.ID),
		Title:       data.Title,
		Description: data.Description,
	}
}
