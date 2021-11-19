package response

import (
	"time"
	"workuo/features/user"
)

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

func toSkillsetResponse(skill user.SkillsetCore) SkillsetResponse {
	return SkillsetResponse{
		Name:     skill.Name,
		Category: skill.Category,
	}
}

func toSkillsetResponseList(skillList []user.SkillsetCore) []SkillsetResponse {
	convertedSkillset := []SkillsetResponse{}

	for _, skill := range skillList {
		convertedSkillset = append(convertedSkillset, toSkillsetResponse(skill))
	}

	return convertedSkillset
}

func toExperienceResponse(experience user.ExperienceCore) ExperienceResponse {
	return ExperienceResponse{
		Title:       experience.Title,
		Description: experience.Description,
		StartDate:   experience.StartDate,
		EndDate:     experience.EndDate,
	}
}

func toExperienceResponseList(experienceList []user.ExperienceCore) []ExperienceResponse {
	convertedExperiences := []ExperienceResponse{}
	for _, exp := range experienceList {
		convertedExperiences = append(convertedExperiences, toExperienceResponse(exp))
	}

	return convertedExperiences
}

func toUserResponse(user user.UserCore) UserResponse {
	return UserResponse{
		Id:          user.Id,
		Name:        user.Name,
		Dob:         user.Dob,
		Gender:      user.Address,
		Address:     user.Address,
		Title:       user.Title,
		Bio:         user.Bio,
		Skillsets:   toSkillsetResponseList(user.Skillsets),
		Experiences: toExperienceResponseList(user.Experiences),
	}
}
