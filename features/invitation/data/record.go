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
	Job         Job
}

type Job struct {
	ID          int
	Title       string
	Description string
	RecruiterId int
}

func ToCore(data Invitation) invitation.InvitationCore {
	return invitation.InvitationCore{
		ID:          data.ID,
		RecruiterID: data.RecruiterID,
		UserID:      data.UserID,
		JobID:       data.JobID,
		Status:      data.Status,
		Job:         ToJobCore(data.Job),
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

func ToJobCore(data Job) invitation.JobCore {
	return invitation.JobCore{
		ID:          data.ID,
		RecruiterId: data.RecruiterId,
		Title:       data.Title,
		Description: data.Description,
	}
}
