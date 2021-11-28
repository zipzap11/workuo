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

func (ir *invitationRepository) GetInvitationByID(id int) (invitation.InvitationCore, error) {
	var data Invitation
	err := ir.DB.Find(&data, id).Error
	if err != nil {
		return invitation.InvitationCore{}, err
	}

	return ToCore(data), nil
}

func (ir *invitationRepository) AcceptInvitation(id int) error {
	err := ir.DB.Model(&Invitation{}).Where("id = ?", id).Update("status", "accepted").Error
	if err != nil {
		return err
	}

	return nil
}

func (ir *invitationRepository) RejectInvitation(id int) error {
	err := ir.DB.Model(&Invitation{}).Where("id = ?", id).Update("status", "rejected").Error
	if err != nil {
		return err
	}

	return nil
}

func (ir *invitationRepository) GetInvitationByUserID(userId int) ([]invitation.InvitationCore, error) {
	var invitations []Invitation
	err := ir.DB.Debug().Where("invitations.user_id = ?", userId).Joins("Job").Find(&invitations).Error
	if err != nil {
		return nil, err
	}

	return ToCoreList(invitations), nil
}
