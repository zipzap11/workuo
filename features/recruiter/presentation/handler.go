package presentation

import (
	"net/http"
	"strconv"
	"workuo/features/recruiter"
	"workuo/features/recruiter/presentation/request"
	"workuo/features/recruiter/presentation/response"

	"github.com/labstack/echo/v4"
)

type RecruiterHandler struct {
	recruiterService recruiter.Service
}

func NewRecruiterHandler(data recruiter.Service) *RecruiterHandler {
	return &RecruiterHandler{data}
}

func (rh *RecruiterHandler) RegisterRecruiterHandler(e echo.Context) error {
	reqData := request.RecruiterRequest{}

	err := e.Bind(&reqData)
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusBadRequest)
	}

	err = rh.recruiterService.RegisterRecruiter(request.FromRecruiterRequest(reqData))
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusInternalServerError)
	}

	return response.NewSuccessResponse(e, "success", nil)
}

func (rp *RecruiterHandler) LoginRecruiterHandler(e echo.Context) error {
	var recruiterLogin request.RecruiterLogin

	err := e.Bind(&recruiterLogin)
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusBadRequest)
	}

	data, err := rp.recruiterService.LoginRecruiter(request.FromRecruiterLogin(recruiterLogin))
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusForbidden)
	}

	return response.NewSuccessResponse(e, "success", response.ToRecruiterLoginResponse(data))
}

func (rh *RecruiterHandler) GetRecruitersHandler(e echo.Context) error {
	data, err := rh.recruiterService.GetRecruiters()
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusInternalServerError)
	}

	return response.NewSuccessResponse(e, "success", response.ToRecruiterResponseList(data))
}

func (rh *RecruiterHandler) GetRecruiterByIdHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusBadRequest)
	}

	data, err := rh.recruiterService.GetRecruiterById(recruiter.RecruiterCore{ID: uint(id)})
	if err != nil {
		return response.NewErrorResponse(e, err.Error(), http.StatusInternalServerError)
	}

	return response.NewSuccessResponse(e, "success", response.ToRecruiterResponse(data))
}
