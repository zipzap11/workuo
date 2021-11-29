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

type UserSkillsetCore struct {
	UserID     uint
	SkillsetId uint
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
	GetUserById(id int) (user UserCore, err error)
	UpdateUser(data UserCore) error
}

type Repository interface {
	InsertUserData(data UserCore) (id int, err error)
	CheckUser(data UserCore) (user UserCore, err error)
	GetData(UserCore) (user []UserCore, err error)
	GetDataById(id int) (user UserCore, err error)
	UpdateUser(data UserCore) error
	CreateExperience(data ExperienceCore) error
	DeleteExperience(id int) error
	UpdateExperience(data ExperienceCore) error
	CreateSkillset(data SkillsetCore) (int, error)
	UpdateUserSkillset(userId int, skillsetId int, newSkillsetId int) error
	DeleteUserSkillset(userId int, skillsetId int) error
	GetUserSkillsets(data UserCore) ([]UserSkillsetCore, error)
	AddUserSkillset(userId int, skillsetId int) error
	GetUserByEmail(email string) (bool, error)
}
