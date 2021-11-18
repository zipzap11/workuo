package request

import (
	"time"
)

type UserRequest struct {
	Name        string              `json: "name"`
	Address     string              `json: "address"`
	Dob         time.Time           `json: "dob"`
	Gender      string              `json: "gender"`
	Bio         string              `json: "bio"`
	Title       string              `json: "title"`
	Skillsets   []SkillsetRequest   `json: "skillsets"`
	Experiences []ExperienceRequest `json: "experiences"`
}

type SkillsetRequest struct {
	Name     string `json: "name"`
	Category string `json: "category"`
}

type ExperienceRequest struct {
	Title       string
	Description string
	StartDate   time.Time
	EndDate     time.Time
}
