package service

import (
	"workuo/features/invitation"
	"workuo/features/job"
	"workuo/features/user"
)

func ToUserCore(data user.UserCore) invitation.UserCore {
	return invitation.UserCore{
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

func ToExperiencesCore(data []user.ExperienceCore) []invitation.ExperienceCore {
	converted := []invitation.ExperienceCore{}
	for _, ex := range data {
		temp := invitation.ExperienceCore{
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

func ToSkillsetsCore(data []user.SkillsetCore) []invitation.SkillsetCore {
	converted := []invitation.SkillsetCore{}
	for _, ex := range data {
		temp := invitation.SkillsetCore{
			Id:       ex.Id,
			Name:     ex.Name,
			Category: ex.Category,
		}
		converted = append(converted, temp)
	}
	return converted
}

func ToJobCore(data job.JobCore) invitation.JobCore {
	return invitation.JobCore{
		ID:           data.ID,
		Title:        data.Title,
		Description:  data.Description,
		RecruiterId:  data.RecruiterId,
		Company:      data.Company,
		Requirements: ToRequirementsCore(data.Requirements),
	}
}

func ToRequirementsCore(data []job.RequirementCore) []invitation.RequirementCore {
	converted := []invitation.RequirementCore{}
	for _, req := range data {
		converted = append(converted, invitation.RequirementCore{
			ID:          req.ID,
			Description: req.Description,
		})
	}
	return converted
}
