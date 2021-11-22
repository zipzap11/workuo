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
	Email       string
	Password    string
	Token       string
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
	LoginUser(data UserCore) (user UserCore, err error)
	GetUsers(data UserCore) (users []UserCore, err error)
	GetUserById(data UserCore) (user UserCore, err error)
	GetUserBySkillsets(data UserCore) (users []UserCore, err error)
}

type Repository interface {
	InsertData(data UserCore) (err error)
	CheckUser(data UserCore) (user UserCore, err error)
	GetData(UserCore) (user []UserCore, err error)
	GetDataById(data UserCore) (user UserCore, err error)
	GetDataBySkillsets(data UserCore) (users []UserCore, err error)
}
