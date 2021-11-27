package presentation

import (
	"net/http"
	"strconv"
	"workuo/features/application"
	"workuo/features/application/presentation/request"
	"workuo/features/application/presentation/response"

	"github.com/labstack/echo/v4"
)

type AppHandler struct {
	appService application.Service
}

func NewAppHandler(as application.Service) *AppHandler {
	return &AppHandler{as}
}

func (ah *AppHandler) ApplyJobHandler(e echo.Context) error {
	var reqPayload request.ApplicationRequest
	err := e.Bind(&reqPayload)
	if err != nil {
		return response.NewSuccessResponse(e, err.Error(), http.StatusBadRequest)
	}

	err = ah.appService.ApplyJob(reqPayload.ToCore())
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusInternalServerError)
	}

	return response.NewSuccessResponse(e, "success", nil)
}

func (ah *AppHandler) GetApplicationByUserIdHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.QueryParam("user-id"))
	if err != nil {
		return response.NewSuccessResponse(e, err.Error(), http.StatusBadRequest)
	}

	applications, err := ah.appService.GetApplicationByUserID(id)
  if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusInternalServerError)
	}
  
  return response.NewSuccessResponse(e, "success", applications)
}

func (ah *AppHandler) RejectApplicationHandler(e echo.Context) error {
  id, err := strconv.Atoi(e.QueryParam("id"))
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusBadRequest)
	}
  
  err = ah.appService.RejectApplication(id)
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
  
	err = ah.appService.AcceptApplication(id)
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusInternalServerError)
	}

	return response.NewSuccessResponse(e, "success", nil)
}
