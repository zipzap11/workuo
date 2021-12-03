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
}

type RequirementCore struct {
	ID          uint
	JobId       uint
	Description string
}

type UserCore struct {
	ID          uint
	Name        string
	Dob         time.Time
	Gender      string
	Address     string
	Title       string
	Bio         string
	Skillsets   []SkillsetCore
	Experiences []ExperienceCore
}

type SkillsetCore struct {
	Id       uint
	Name     string
	Category string
}

type ExperienceCore struct {
	Id          uint
	UserId      uint
	Description string
	Title       string
	StartDate   time.Time
	EndDate     time.Time
}

type Repository interface {
	ApplyJob(ApplicationCore) error
	GetApplicationByUserID(int) ([]ApplicationCore, error)
	GetApplicationByID(int) (ApplicationCore, error)
	GetApplicationByJobID(int) ([]ApplicationCore, error)
	GetApplicationMultiParam(int, int) (ApplicationCore, error)
	RejectApplication(int) error
	AcceptApplication(int) error
}

type Service interface {
	ApplyJob(ApplicationCore) error
	GetApplicationByUserID(int) ([]ApplicationCore, error)
	GetApplicationByID(int) (ApplicationCore, error)
	GetApplicationByJobID(int) ([]ApplicationCore, error)
	GetApplicationMultiParam(int, int) (ApplicationCore, error)
	RejectApplication(int, int) error
	AcceptApplication(int, int) error
}
