package data

import (
	"time"

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
	Category uint
}

type Experience struct {
	ID          uint
	UserID      uint
	Description string
	Title       string
	StartDate   time.Time
	EndDate     time.Time
}
