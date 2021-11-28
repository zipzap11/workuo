package presentation

import (
	"net/http"
	"strconv"
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

func (ih *InvitationHandler) GetInvitationByIDHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return response.NewErrorResponse(e, http.StatusBadRequest, err.Error())
	}

	data, err := ih.invService.GetInvitationByID(id)
	if err != nil {
		return response.NewErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return response.NewSuccessResponse(e, response.ToInvitationResponse(data))
}

func (ih *InvitationHandler) AcceptInvitation(e echo.Context) error {
	id, err := strconv.Atoi(e.QueryParam("id"))
	if err != nil {
		return response.NewErrorResponse(e, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractClaim(e)
	role := claims["role"]
	userId := int(claims["id"].(float64))
	if role != "user" {
		return response.NewErrorResponse(e, http.StatusBadRequest, "role not allowed to accept invitation")
	}

	err = ih.invService.AcceptInvitation(userId, id)
	if err != nil {
		return response.NewErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return response.NewSuccessResponse(e, nil)
}
