package presentation

import (
	"errors"
	"net/http"
	"strconv"
	"workuo/features/job"
	"workuo/features/job/presentation/request"
	"workuo/features/job/presentation/response"
	"workuo/helper"
	"workuo/middleware"

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
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid payload data", err)
	}

	claims := middleware.ExtractClaim(e)
	role := claims["role"].(string)
	if role != "recruiter" {
		return helper.ErrorResponse(e, http.StatusForbidden, "user not allowed to create job post", errors.New("not allowed"))
	}

	recId := claims["id"].(float64)
	payloadData.RecruiterId = int(recId)

	err = jh.jobService.CreateJobPost(payloadData.ToCore())
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}

	return helper.SuccessResponse(e, nil)
}

func (jh *JobHandler) GetJobPostHandler(e echo.Context) error {
	var reqData request.JobFilter
	err := e.Bind(&reqData)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid payload data", err)
	}

	data, err := jh.jobService.GetJobPost(reqData.ToCore())
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}

	return helper.SuccessResponse(e, response.ToJobResponseList(data))
}

func (jh *JobHandler) GetJobPostByIdHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid id parameter", err)
	}
	data, err := jh.jobService.GetJobPostById(id)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}

	return helper.SuccessResponse(e, response.ToJobResponse(data))
}

func (jh *JobHandler) DeleteJobPostHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid id parameter", err)
	}

	claims := middleware.ExtractClaim(e)
	recId := claims["id"].(float64)
	role := claims["role"].(string)
	if role != "recruiter" {
		return helper.ErrorResponse(e, http.StatusForbidden, "role not allowed to delete data", errors.New("not allowed"))
	}

	err = jh.jobService.DeleteJobPost(job.JobCore{ID: id, RecruiterId: int(recId)})
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}

	return helper.SuccessResponse(e, nil)
}

func (jh *JobHandler) UpdateJobPostHandler(e echo.Context) error {
	payloadData := request.JobUpdate{}

	err := e.Bind(&payloadData)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid payload data", err)
	}

	claims := middleware.ExtractClaim(e)

	payloadData.RecruiterId = int(claims["id"].(float64))
	role := claims["role"].(string)
	if role != "recruiter" {
		return helper.ErrorResponse(e, http.StatusForbidden, "role not allowed to delete data", errors.New("not allowed"))
	}

	err = jh.jobService.UpdateJobPost(payloadData.ToCore())
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}

	return helper.SuccessResponse(e, nil)
}
