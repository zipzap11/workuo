package service

import "workuo/features/invitation"

type invitationService struct {
	invRepository invitation.Repository
}

func NewInvitationService(ir invitation.Repository) invitation.Service {
	return &invitationService{ir}
}
