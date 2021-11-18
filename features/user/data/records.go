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
	Skillsets   []Skillset `gorm: "many2many:user_skillsets;"`
	Experiences []Experience
}

type Skillset struct {
	ID       uint
	UserID   uint
	Category string
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
			Category: s.Category,
		})
	}

	return convertedSkillsets
}
