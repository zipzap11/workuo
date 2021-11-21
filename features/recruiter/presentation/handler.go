package presentation

import (
	"net/http"
	"workuo/features/recruiter"
	"workuo/features/recruiter/presentation/request"

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

	err := e.Bind(reqData)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	err = rh.recruiterService.RegisterRecruiter(request.ToCore(reqData))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}
