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

type Skillset struct {
	ID       uint
	Category string
	Name     string
}

type Experience struct {
	ID          uint
	UserID      uint
	Description string
	Title       string
	StartDate   time.Time
	EndDate     time.Time
}

func toSkillsetRecords(skillsets []user.SkillsetCore) []Skillset {
	convertedSkillsets := []Skillset{}
	for _, s := range skillsets {
		convertedSkillsets = append(convertedSkillsets, Skillset{
			Name:     s.Name,
			Category: s.Category,
		})
	}

	return convertedSkillsets
}

func toExperienceRecords(experiences []user.ExperienceCore) []Experience {
	convertedExperiences := []Experience{}
	for _, ex := range experiences {
		convertedExperiences = append(convertedExperiences, Experience{
			Title:       ex.Title,
			Description: ex.Description,
			StartDate:   ex.StartDate,
			EndDate:     ex.EndDate,
		})
	}

	return convertedExperiences
}

func toUserRecord(user user.UserCore) User {
	return User{
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
