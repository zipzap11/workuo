package invitation

import "time"

type InvitationCore struct {
	ID          uint
	RecruiterID uint
	UserID      uint
	JobID       uint
	Status      string
	Role        string
	User        UserCore
	Job         JobCore
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

type Service interface {
	InviteUser(InvitationCore) error
	GetInvitationByID(id int) (InvitationCore, error)
}

type Repository interface {
	InviteUser(InvitationCore) error
	GetInvitationByID(id int) (InvitationCore, error)
}
