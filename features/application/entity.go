package application

import (
	"time"
)

type ApplicationCore struct {
	ID        uint
	UserID    uint
	JobID     uint
	Status    string
	AppliedAt time.Time
	Job       JobCore
}

type JobCore struct {
	ID           int
	Title        string
	Description  string
	RecruiterId  int
	Company      string
	Requirements []RequirementCore
	Created_at   time.Time
	Updated_at   time.Time
}

type RequirementCore struct {
	ID          uint
	JobId       uint
	Description string
}

type Repository interface {
	ApplyJob(ApplicationCore) error
	GetApplicationByUserID(int) ([]ApplicationCore, error)
	GetApplicationByID(int) (ApplicationCore, error)
	RejectApplication(int) error
	AcceptApplication(int) error
}

type Service interface {
	ApplyJob(ApplicationCore) error
	GetApplicationByUserID(int) ([]ApplicationCore, error)
	GetApplicationByID(int) (ApplicationCore, error)
	RejectApplication(int) error
	AcceptApplication(int) error
}
