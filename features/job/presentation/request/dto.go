package request

import "workuo/features/job"

type Job struct {
	Title        string   `json: "title"`
	Description  string   `json: "description"`
	RecruiterId  int      `json: "recruiterId"`
	Requirements []string `json: "requirements"`
}

func (j *Job) ToCore() job.JobCore {
	convertedRequirement := []job.RequirementCore{}
	for _, req := range j.Requirements {
		convertedRequirement = append(convertedRequirement, job.RequirementCore{
			Description: req,
		})
	}

	return job.JobCore{
		Title:        j.Title,
		RecruiterId:  j.RecruiterId,
		Description:  j.Description,
		Requirements: convertedRequirement,
	}
}
