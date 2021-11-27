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
	User      UserCore
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

type UserCore struct {
	ID      uint
	Name    string
	Dob     time.Time
	Gender  string
	Address string
	Title   string
	Bio     string
}

type Repository interface {
	ApplyJob(ApplicationCore) error
	GetApplicationByUserID(int) ([]ApplicationCore, error)
	GetApplicationByID(int) (ApplicationCore, error)
	GetApplicationByJobID(int) ([]ApplicationCore, error)
	RejectApplication(int) error
	AcceptApplication(int) error
}

type Service interface {
	ApplyJob(ApplicationCore) error
	GetApplicationByUserID(int) ([]ApplicationCore, error)
	GetApplicationByID(int) (ApplicationCore, error)
	GetApplicationByJobID(int) ([]ApplicationCore, error)
	RejectApplication(int, int) error
	AcceptApplication(int, int) error
}
