package presentation

import (
	"net/http"
	"workuo/features/invitation"
	"workuo/features/invitation/presentation/request"
	"workuo/features/invitation/presentation/response"
	"workuo/middleware"

	"github.com/labstack/echo/v4"
)

type InvitationHandler struct {
	invService invitation.Service
}

func NewInvitationHandler(is invitation.Service) *InvitationHandler {
	return &InvitationHandler{is}
}

func (ih *InvitationHandler) InviteUserHandler(e echo.Context) error {
	var payloadData request.InvitationRequest
	err := e.Bind(&payloadData)
	if err != nil {
		return response.NewErrorResponse(e, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractClaim(e)
	payloadData.RecruiterID = uint(claims["id"].(float64))
	payloadData.Role = claims["role"].(string)

	err = ih.invService.InviteUser(request.ToCore(payloadData))
	if err != nil {
		return response.NewErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return response.NewSuccessResponse(e, nil)
}
