package request

import (
	"time"
	"workuo/features/user"
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

func (sr *SkillsetRequest) toSkillsetCore() user.SkillsetCore {
	return user.SkillsetCore{
		Name:     sr.Name,
		Category: sr.Category,
	}
}

func toSkillsetsCore(srs []SkillsetRequest) []user.SkillsetCore {
	convertedSkillset := []user.SkillsetCore{}
	for _, skill := range srs {
		convertedSkillset = append(convertedSkillset, skill.toSkillsetCore())
	}

	return convertedSkillset
}

func (ec *ExperienceRequest) toExperienceCore() user.ExperienceCore {
	return user.ExperienceCore{
		Title:       ec.Title,
		Description: ec.Description,
		StartDate:   ec.StartDate,
		EndDate:     ec.EndDate,
	}
}

func toExperiencesCore(ex []ExperienceRequest) []user.ExperienceCore {
	convertedExperiences := []user.ExperienceCore{}
	for _, exp := range ex {
		convertedExperiences = append(convertedExperiences, exp.toExperienceCore())
	}

	return convertedExperiences
}

func (ur *UserRequest) toUserCore() user.UserCore {
	return user.UserCore{
		Name:        ur.Name,
		Address:     ur.Address,
		Dob:         ur.Dob,
		Gender:      ur.Gender,
		Bio:         ur.Bio,
		Title:       ur.Title,
		Skillsets:   toSkillsetsCore(ur.Skillsets),
		Experiences: toExperiencesCore(ur.Experiences),
	}
}
