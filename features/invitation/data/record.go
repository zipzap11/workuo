package data

import (
	"workuo/features/invitation"

	"gorm.io/gorm"
)

type Invitation struct {
	gorm.Model
	ID          uint
	RecruiterID uint
	UserID      uint
	JobID       uint
	Status      string
}

func ToCore(data Invitation) invitation.InvitationCore {
	return invitation.InvitationCore{
		ID:          data.ID,
		RecruiterID: data.RecruiterID,
		UserID:      data.UserID,
		JobID:       data.JobID,
		Status:      data.Status,
	}
}

func FromCore(data invitation.InvitationCore) Invitation {
	return Invitation{
		ID:          data.ID,
		RecruiterID: data.RecruiterID,
		UserID:      data.UserID,
		JobID:       data.JobID,
		Status:      data.Status,
	}
}

func ToCoreList(data []Invitation) []invitation.InvitationCore {
	convertedData := []invitation.InvitationCore{}
	for _, data := range data {
		convertedData = append(convertedData, ToCore(data))
	}
	return convertedData
}
