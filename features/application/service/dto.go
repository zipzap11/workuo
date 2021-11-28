package service

import (
	"workuo/features/application"
	"workuo/features/job"
	"workuo/features/user"
)

func ToUserCore(data user.UserCore) application.UserCore {
	return application.UserCore{
		ID:          data.Id,
		Name:        data.Name,
		Dob:         data.Dob,
		Gender:      data.Gender,
		Address:     data.Address,
		Title:       data.Title,
		Bio:         data.Bio,
		Skillsets:   ToSkillsetsCore(data.Skillsets),
		Experiences: ToExperiencesCore(data.Experiences),
	}
}

func ToExperiencesCore(data []user.ExperienceCore) []application.ExperienceCore {
	converted := []application.ExperienceCore{}
	for _, ex := range data {
		temp := application.ExperienceCore{
			Id:          ex.Id,
			Title:       ex.Title,
			Description: ex.Description,
			StartDate:   ex.StartDate,
			EndDate:     ex.EndDate,
		}
		converted = append(converted, temp)
	}
	return converted
}

func ToSkillsetsCore(data []user.SkillsetCore) []application.SkillsetCore {
	converted := []application.SkillsetCore{}
	for _, ex := range data {
		temp := application.SkillsetCore{
			Id:       ex.Id,
			Name:     ex.Name,
			Category: ex.Category,
		}
		converted = append(converted, temp)
	}
	return converted
}

func ToJobCore(data job.JobCore) application.JobCore {
	return application.JobCore{
		ID:           data.ID,
		Title:        data.Title,
		Description:  data.Description,
		RecruiterId:  data.RecruiterId,
		Company:      data.Company,
		Requirements: ToRequirementsCore(data.Requirements),
	}
}

func ToRequirementsCore(data []job.RequirementCore) []application.RequirementCore {
	converted := []application.RequirementCore{}
	for _, req := range data {
		converted = append(converted, application.RequirementCore{
			ID:          req.ID,
			Description: req.Description,
		})
	}
	return converted
}
