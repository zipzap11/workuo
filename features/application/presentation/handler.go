package presentation

import (
	"net/http"
	"strconv"
	"workuo/features/application"
	"workuo/features/application/presentation/response"
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
		return response.NewSuccessResponse(e, err.Error(), http.StatusBadRequest)
	}

	claims := middleware.ExtractClaim(e)
	role := claims["role"]
	userId := uint(claims["id"].(float64))
	if role != "user" {
		return response.NewErrorResponse(e, "only user role can apply job", http.StatusForbidden)
	}

	err = ah.appService.ApplyJob(application.ApplicationCore{
		JobID:  uint(jobId),
		UserID: userId,
	})
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusInternalServerError)
	}

	return response.NewSuccessResponse(e, "success", nil)
}

func (ah *AppHandler) GetApplicationByUserIdHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return response.NewSuccessResponse(e, err.Error(), http.StatusBadRequest)
	}

	applications, err := ah.appService.GetApplicationByUserID(id)
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusInternalServerError)
	}

	return response.NewSuccessResponse(e, "success", response.ToApplicationResponseUserList(applications))
}

func (ah *AppHandler) RejectApplicationHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.QueryParam("id"))
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusBadRequest)
	}

	claims := middleware.ExtractClaim(e)
	role := claims["role"]
	recruiterId := uint(claims["id"].(float64))
	if role != "recruiter" {
		return response.NewErrorResponse(e, "user not allowed to reject application", http.StatusForbidden)
	}

	err = ah.appService.RejectApplication(id, int(recruiterId))
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusInternalServerError)
	}

	return response.NewSuccessResponse(e, "success", nil)
}

func (ah *AppHandler) AcceptApplication(e echo.Context) error {
	id, err := strconv.Atoi(e.QueryParam("id"))
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusBadRequest)
	}

	claims := middleware.ExtractClaim(e)
	role := claims["role"]
	recruiterId := uint(claims["id"].(float64))
	if role != "recruiter" {
		return response.NewErrorResponse(e, "user not allowed to accept application", http.StatusForbidden)
	}

	err = ah.appService.AcceptApplication(id, int(recruiterId))
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusInternalServerError)
	}

	return response.NewSuccessResponse(e, "success", nil)
}

func (ah *AppHandler) GetApplicationByIDHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusBadRequest)
	}

	appCore, err := ah.appService.GetApplicationByID(id)
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusInternalServerError)
	}

	return response.NewSuccessResponse(e, "success", response.ToApplicationResponse(appCore))
}

func (ah *AppHandler) GetApplicationByJobIDHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusBadRequest)
	}

	apps, err := ah.appService.GetApplicationByJobID(id)
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusInternalServerError)
	}

	return response.NewSuccessResponse(e, "success", response.ToApplicationResponseJobList(apps))
}
