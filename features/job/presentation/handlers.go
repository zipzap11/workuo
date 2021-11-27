package presentation

import (
	"net/http"
	"strconv"
	"workuo/features/job"
	"workuo/features/job/presentation/request"
	"workuo/features/job/presentation/response"

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

	err = jh.jobService.CreateJobPost(payloadData.ToCore())

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

func (jh *JobHandler) GetJobPostHandler(e echo.Context) error {
	var reqData request.JobFilter
	err := e.Bind(&reqData)
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusBadRequest)
	}

	data, err := jh.jobService.GetJobPost(reqData.ToCore())
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusInternalServerError)
	}

	return response.NewSuccessResponse(e, "success", response.ToJobResponseList(data))
}

func (jh *JobHandler) GetJobPostByIdHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusBadRequest)
	}
	data, err := jh.jobService.GetJobPostById(id)
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusInternalServerError)
	}

	return response.NewSuccessResponse(e, "success", response.ToJobResponse(data))
}


func (jh *JobHandler) DeleteJobPostHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
  if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusBadRequest)
	}
  
  err = jh.jobService.DeleteJobPost(job.JobCore{ID: id})
  if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusInternalServerError)
	}
  
  return response.NewSuccessResponse(e, "success", nil)
}

func (jh *JobHandler) UpdateJobPostHandler(e echo.Context) error {
	payloadData := request.JobUpdate{}

	err := e.Bind(&payloadData)
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusBadRequest)
	}
  
  err = jh.jobService.UpdateJobPost(payloadData.ToCore())
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusInternalServerError)
	}

	return response.NewSuccessResponse(e, "success", nil)
}
