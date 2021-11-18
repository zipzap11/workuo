package request

import "workuo/features/job"

type Job struct {
	Title        string   `json: "title"`
	Description  string   `json: "description"`
	RecruiterID  int      `json: "recruiter_id"`
	Requirements []string `json: "requirements"`
}

func (j *Job) toCore() job.JobCore {
	convertedRequirement := []job.RequirementCore{}
	for _, req := range j.Requirements {
		convertedRequirement = append(convertedRequirement, job.RequirementCore{
			Description: req,
		})
	}

	return job.JobCore{
		Title:        j.Title,
		Recruiter_id: j.RecruiterID,
		Description:  j.Description,
		Requirements: convertedRequirement,
	}
}
