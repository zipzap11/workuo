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

func (ir *invitationRepository) InviteUser(data invitation.InvitationCore) error {
	recordData := FromCore(data)
	err := ir.DB.Create(&recordData).Error
	if err != nil {
		return err
	}

	return nil
}
