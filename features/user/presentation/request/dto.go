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
	Email       string              `json: "email"`
	Password    string              `json: "password"`
	Skillsets   []SkillsetRequest   `json: "skillsets"`
	Experiences []ExperienceRequest `json: "experiences"`
}

type SkillsetRequest struct {
	Name     string `json: "name"`
	Category string `json: "category"`
}

type ExperienceRequest struct {
	Title       string    `json: "title"`
	Description string    `json: "description"`
	StartDate   time.Time `json: "startDate"`
	EndDate     time.Time `json: "endDate"`
}

type UserAuth struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

type UserFilter struct {
	Title     string            `json: "title"`
	Skillsets []SkillsetRequest `json: "skillsets`
}

func (data *UserAuth) ToUserCore() user.UserCore {
	return user.UserCore{
		Email:    data.Email,
		Password: data.Password,
	}
}

func (sr *SkillsetRequest) toSkillsetCore() user.SkillsetCore {
	return user.SkillsetCore{
		Name:     sr.Name,
		Category: sr.Category,
	}
}

func ToSkillsetsCore(srs []SkillsetRequest) []user.SkillsetCore {
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

func (requestData *UserRequest) ToUserCore() user.UserCore {
	return user.UserCore{
		Name:        requestData.Name,
		Address:     requestData.Address,
		Dob:         requestData.Dob,
		Gender:      requestData.Gender,
		Bio:         requestData.Bio,
		Title:       requestData.Title,
		Email:       requestData.Email,
		Password:    requestData.Password,
		Skillsets:   ToSkillsetsCore(requestData.Skillsets),
		Experiences: toExperiencesCore(requestData.Experiences),
	}
}
