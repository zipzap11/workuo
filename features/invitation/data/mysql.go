package data

import (
	"workuo/features/invitation"

	"gorm.io/gorm"
)

type invitationRepository struct {
	DB *gorm.DB
}

func NewInvitationRepository(DB *gorm.DB) invitation.Repository {
	return &invitationRepository{DB}
}
