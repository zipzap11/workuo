package response

import "time"

type UserResponse struct {
	Id          uint                 `json: "id"`
	Name        string               `json: "name"`
	Dob         time.Time            `json: "dob"`
	Gender      string               `json: "gender"`
	Address     string               `json: "string"`
	Title       string               `json: "title"`
	Bio         string               `json: "bio"`
	Skillsets   []SkillsetResponse   `json: "skillsets`
	Experiences []ExperienceResponse `json: "experiences"`
}

type SkillsetResponse struct {
	Name     string `json: "name"`
	Category string `json: "category"`
}

type ExperienceResponse struct {
	Title       string    `json: "title"`
	Description string    `json: "description"`
	StartDate   time.Time `json: "startDate"`
	EndDate     time.Time `json: "startDate"`
}
