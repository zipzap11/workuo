package data

import (
	"time"
	"workuo/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string
	Dob         time.Time
	Gender      string
	Address     string
	Title       string
	Bio         string
	Email       string
	Password    string
	Skillsets   []Skillset `gorm:"many2many:user_skillsets;"`
	Experiences []Experience
}

type UserSkillset struct {
	UserID     uint `gorm: "primaryKey"`
	SkillsetID uint `gorm: "primaryKey"`
}

type Skillset struct {
	ID       uint
	Category string
	Name     string `gorm: "unique;"`
}

type Experience struct {
	ID          uint
	UserID      uint
	Description string
	Title       string
	StartDate   time.Time
	EndDate     time.Time
}

func ToUserSkillsetCore(data []UserSkillset) []user.UserSkillsetCore {
	converted := []user.UserSkillsetCore{}
	for _, skill := range data {
		converted = append(converted, user.UserSkillsetCore{
			UserID:     skill.UserID,
			SkillsetId: skill.SkillsetID,
		})
	}
	return converted
}

func toSkillsetRecords(skillsets []user.SkillsetCore) []Skillset {
	convertedSkillsets := []Skillset{}
	for _, s := range skillsets {
		convertedSkillsets = append(convertedSkillsets, ToSkillsetRecord(s))
	}
	return convertedSkillsets
}

func ToSkillsetRecord(data user.SkillsetCore) Skillset {
	return Skillset{
		ID:       data.Id,
		Name:     data.Name,
		Category: data.Category,
	}
}

func toExperienceRecords(experiences []user.ExperienceCore) []Experience {
	convertedExperiences := []Experience{}
	for _, ex := range experiences {
		convertedExperiences = append(convertedExperiences, ToExperienceRecord(ex))
	}

	return convertedExperiences
}

func ToExperienceRecord(data user.ExperienceCore) Experience {
	return Experience{
		ID:          data.Id,
		UserID:      data.UserId,
		Title:       data.Title,
		Description: data.Description,
		StartDate:   data.StartDate,
		EndDate:     data.EndDate,
	}
}

func toUserRecord(user user.UserCore) User {
	return User{
		Model: gorm.Model{
			ID: user.Id,
		},
		Name:        user.Name,
		Dob:         user.Dob,
		Gender:      user.Gender,
		Address:     user.Address,
		Title:       user.Title,
		Bio:         user.Bio,
		Email:       user.Email,
		Password:    user.Password,
		Skillsets:   toSkillsetRecords(user.Skillsets),
		Experiences: toExperienceRecords(user.Experiences),
	}
}

func toExperienceCore(ex Experience) user.ExperienceCore {
	return user.ExperienceCore{
		Id:          ex.ID,
		UserId:      ex.UserID,
		Title:       ex.Title,
		Description: ex.Description,
		StartDate:   ex.StartDate,
		EndDate:     ex.EndDate,
	}
}

func toExperienceCoreList(expList []Experience) []user.ExperienceCore {
	convertedExperience := []user.ExperienceCore{}
	for _, ex := range expList {
		convertedExperience = append(convertedExperience, toExperienceCore(ex))
	}

	return convertedExperience
}

func toSkillsetCore(skill Skillset) user.SkillsetCore {
	return user.SkillsetCore{
		Id:       skill.ID,
		Name:     skill.Name,
		Category: skill.Category,
	}
}

func toSkillsetCoreList(skillList []Skillset) []user.SkillsetCore {
	convertedSkillsets := []user.SkillsetCore{}

	for _, skill := range skillList {
		convertedSkillsets = append(convertedSkillsets, toSkillsetCore(skill))
	}

	return convertedSkillsets
}

func toUserCore(u User) user.UserCore {
	return user.UserCore{
		Id:          u.ID,
		Name:        u.Name,
		Dob:         u.Dob,
		Gender:      u.Gender,
		Title:       u.Title,
		Address:     u.Address,
		Bio:         u.Bio,
		Skillsets:   toSkillsetCoreList(u.Skillsets),
		Experiences: toExperienceCoreList(u.Experiences),
	}
}

func toUserCoreList(uList []User) []user.UserCore {
	convertedUser := []user.UserCore{}

	for _, user := range uList {
		convertedUser = append(convertedUser, toUserCore(user))
	}

	return convertedUser
}

func SeparateUserData(data user.UserCore) (User, []Skillset, []Experience) {
	user := User{
		Name:     data.Name,
		Dob:      data.Dob,
		Gender:   data.Gender,
		Address:  data.Address,
		Title:    data.Title,
		Bio:      data.Bio,
		Email:    data.Email,
		Password: data.Password,
	}

	return user, user.Skillsets, user.Experiences
}
