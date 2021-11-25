package presentation

import "workuo/features/invitation"

type InvitationHandler struct {
	invService invitation.Service
}

func NewInvitationHandler(is invitation.Service) *InvitationHandler {
	return &InvitationHandler{is}
}
