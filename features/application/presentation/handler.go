package presentation

import (
	"errors"
	"net/http"
	"strconv"
	"workuo/features/application"
	"workuo/features/application/presentation/response"
	"workuo/helper"
	"workuo/middleware"

	"github.com/labstack/echo/v4"
)

type AppHandler struct {
	appService application.Service
}

func NewAppHandler(as application.Service) *AppHandler {
	return &AppHandler{as}
}

func (ah *AppHandler) ApplyJobHandler(e echo.Context) error {
	jobId, err := strconv.Atoi(e.QueryParam("jobId"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "Invalid id paramter", err)
	}

	claims := middleware.ExtractClaim(e)
	role := claims["role"]
	userId := uint(claims["id"].(float64))
	if role != "user" {
		return helper.ErrorResponse(e, http.StatusForbidden, "only user role can apply job", err)
	}

	err = ah.appService.ApplyJob(application.ApplicationCore{
		JobID:  uint(jobId),
		UserID: userId,
	})
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "Something went wrong", err)
	}

	return helper.SuccessResponse(e, nil)
}

func (ah *AppHandler) GetApplicationByUserIdHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid id parameter", err)
	}

	applications, err := ah.appService.GetApplicationByUserID(id)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "Something went wrong", err)
	}

	return helper.SuccessResponse(e, response.ToApplicationResponseUserList(applications))
}

func (ah *AppHandler) RejectApplicationHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.QueryParam("id"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid id parameter", err)
	}

	claims := middleware.ExtractClaim(e)
	role := claims["role"]
	recruiterId := uint(claims["id"].(float64))
	if role != "recruiter" {
		return helper.ErrorResponse(e, http.StatusForbidden, "user not allowed to reject application", errors.New("action not allowed"))
	}

	err = ah.appService.RejectApplication(id, int(recruiterId))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}

	return response.NewSuccessResponse(e, "success", nil)
}

func (ah *AppHandler) AcceptApplication(e echo.Context) error {
	id, err := strconv.Atoi(e.QueryParam("id"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid id parameter", err)
	}

	claims := middleware.ExtractClaim(e)
	role := claims["role"]
	recruiterId := uint(claims["id"].(float64))
	if role != "recruiter" {
		return helper.ErrorResponse(e, http.StatusForbidden, "user not allowed to accept application", errors.New("action not allowed"))
	}

	err = ah.appService.AcceptApplication(id, int(recruiterId))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}

	return helper.SuccessResponse(e, nil)
}

func (ah *AppHandler) GetApplicationByIDHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid id parameter", err)
	}

	appCore, err := ah.appService.GetApplicationByID(id)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}

	return helper.SuccessResponse(e, response.ToApplicationResponse(appCore))
}

func (ah *AppHandler) GetApplicationByJobIDHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid id paramter", err)
	}

	apps, err := ah.appService.GetApplicationByJobID(id)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}

	return helper.SuccessResponse(e, response.ToApplicationResponseJobList(apps))
}
