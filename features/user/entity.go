package user

import "time"

type UserCore struct {
	Id          uint
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
	RegisterUser(data UserCore) (err error)
	GetAllUser() (users []UserCore, err error)
}

type Repository interface {
	InsertData(data UserCore) (err error)
	GetData() (user []UserCore, err error)
}
