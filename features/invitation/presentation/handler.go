package presentation

import (
	"errors"
	"net/http"
	"strconv"
	"workuo/features/invitation"
	"workuo/features/invitation/presentation/request"
	"workuo/features/invitation/presentation/response"
	"workuo/helper"
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
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid payload data", err)
	}

	claims := middleware.ExtractClaim(e)
	payloadData.RecruiterID = uint(claims["id"].(float64))
	payloadData.Role = claims["role"].(string)

	err = ih.invService.InviteUser(request.ToCore(payloadData))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}

	return helper.SuccessResponse(e, nil)
}

func (ih *InvitationHandler) GetInvitationByIDHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid id parameter", err)
	}

	data, err := ih.invService.GetInvitationByID(id)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}

	return helper.SuccessResponse(e, response.ToInvitationDetailResponse(data))
}

func (ih *InvitationHandler) AcceptInvitationHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.QueryParam("id"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid id parameter", err)
	}

	claims := middleware.ExtractClaim(e)
	role := claims["role"]
	userId := int(claims["id"].(float64))
	if role != "user" {
		return helper.ErrorResponse(e, http.StatusBadRequest, "role not allowed to accept invitation", errors.New("not allowed"))
	}

	err = ih.invService.AcceptInvitation(userId, id)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}

	return helper.SuccessResponse(e, nil)
}

func (ih *InvitationHandler) RejectInvitationHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.QueryParam("id"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid id parameter", err)
	}

	claims := middleware.ExtractClaim(e)
	role := claims["role"]
	userId := int(claims["id"].(float64))
	if role != "user" {
		return helper.ErrorResponse(e, http.StatusBadRequest, "role not allowed to reject invitation", errors.New("not allowed"))
	}

	err = ih.invService.RejectInvitation(userId, id)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}

	return helper.SuccessResponse(e, nil)
}

func (ih *InvitationHandler) GetInvitationByUserID(e echo.Context) error {
	claims := middleware.ExtractClaim(e)
	role := claims["role"]
	userId := int(claims["id"].(float64))
	if role != "user" {
		return helper.ErrorResponse(e, http.StatusBadRequest, "role not allowed to get invitations", errors.New("not allowed"))
	}

	data, err := ih.invService.GetInvitationByUserID(userId)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}

	return helper.SuccessResponse(e, response.ToInvitationUserResponseList(data))
}

func (ih *InvitationHandler) GetInvitationByJobID(e echo.Context) error {
	jobId, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "role not allowed to get invitations", errors.New("not allowed"))
	}

	claims := middleware.ExtractClaim(e)
	recId := int(claims["id"].(float64))
	role := claims["role"]
	if role != "recruiter" {
		return helper.ErrorResponse(e, http.StatusBadRequest, "role not allowed to get invitations", errors.New("not allowed"))
	}

	data, err := ih.invService.GetInvitationByJobID(jobId, recId)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}

	return helper.SuccessResponse(e, response.ToInvitationJobResponseList(data))
}
