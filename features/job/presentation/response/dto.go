package response

import (
	"net/http"
	"workuo/features/job"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string
	Data    interface{}
}

type JobResponse struct {
	ID           uint
	Title        string
	Description  string
	RecruiterId  int
	Requirements []string
}

func NewSuccessResponse(e echo.Context, msg string, data interface{}) error {
	return e.JSON(http.StatusOK, Response{
		Message: msg,
		Data:    data,
	})
}

func NewErrorResponse(e echo.Context, msg string, code int) error {
	return e.JSON(code, Response{
		Message: msg,
	})
}

func ToJobResponse(data job.JobCore) JobResponse {
	convertedRecquirements := []string{}
	for _, req := range data.Requirements {
		convertedRecquirements = append(convertedRecquirements, req.Description)
	}

	return JobResponse{
		ID:           uint(data.ID),
		Title:        data.Title,
		Description:  data.Description,
		RecruiterId:  data.RecruiterId,
		Requirements: convertedRecquirements,
	}
}

func ToJobResponseList(data []job.JobCore) []JobResponse {
	convertedJob := []JobResponse{}
	for _, job := range data {
		convertedJob = append(convertedJob, ToJobResponse(job))
	}

	return convertedJob
}
