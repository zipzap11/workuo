package data

import (
	"time"
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
	User        User
}

type Job struct {
	ID          int
	Title       string
	Description string
	RecruiterId int
}

type User struct {
	ID      int
	Name    string
	Dob     time.Time
	Gender  string
	Address string
	Title   string
	Bio     string
}

func ToCore(data Invitation) invitation.InvitationCore {
	return invitation.InvitationCore{
		ID:          data.ID,
		RecruiterID: data.RecruiterID,
		UserID:      data.UserID,
		JobID:       data.JobID,
		Status:      data.Status,
		Job:         ToJobCore(data.Job),
		User:        ToUserCore(data.User),
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

func ToUserCore(data User) invitation.UserCore {
	return invitation.UserCore{
		ID:      uint(data.ID),
		Name:    data.Name,
		Title:   data.Title,
		Dob:     data.Dob,
		Gender:  data.Gender,
		Address: data.Address,
		Bio:     data.Bio,
	}
}
