package presentation

import (
	"net/http"
	"workuo/features/job"
	"workuo/features/job/presentation/request"

	"github.com/labstack/echo/v4"
)

type JobHandler struct {
	jobService job.Service
}

func NewJobHandler(js job.Service) *JobHandler {
	return &JobHandler{js}
}

func (jh *JobHandler) CreateJobPostHandler(e echo.Context) error {
	payloadData := request.Job{}
	err := e.Bind(&payloadData)

	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	err = jh.jobService.CreateJobPost(payloadData)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}
