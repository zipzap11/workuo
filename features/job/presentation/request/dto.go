package request

import "workuo/features/job"

type Job struct {
	ID           uint     `json: "id"`
	Title        string   `json: "title"`
	Description  string   `json: "description"`
	RecruiterId  int      `json: "recruiterId"`
	Requirements []string `json: "requirements"`
}

type JobUpdate struct {
	ID           uint          `json: "id"`
	Title        string        `json: "title"`
	Description  string        `json: "description"`
	RecruiterId  int           `json: "recruiterId"`
	Requirements []Requirement `json: "requirements"`
}

type JobFilter struct {
	Title   string `json: "title"`
	Company string `json: "company"`
}

type Requirement struct {
	ID          uint   `json: "id"`
	Description string `json: "description"`
}

func (j *Job) ToCore() job.JobCore {
	convertedRequirement := []job.RequirementCore{}
	for _, req := range j.Requirements {
		convertedRequirement = append(convertedRequirement, job.RequirementCore{
			Description: req,
		})
	}

	return job.JobCore{
		ID:           int(j.ID),
		Title:        j.Title,
		RecruiterId:  j.RecruiterId,
		Description:  j.Description,
		Requirements: convertedRequirement,
	}
}

func (jf *JobFilter) ToCore() job.JobCore {
	return job.JobCore{
		Title:   jf.Title,
		Company: jf.Company,
	}
}

func (r *Requirement) ToCore() job.RequirementCore {
	return job.RequirementCore{
		ID:          r.ID,
		Description: r.Description,
	}
}

func (ju *JobUpdate) ToCore() job.JobCore {
	convertedReqs := []job.RequirementCore{}
	for _, req := range ju.Requirements {
		convertedReqs = append(convertedReqs, req.ToCore())
	}

	return job.JobCore{
		ID:           int(ju.ID),
		Title:        ju.Title,
		Description:  ju.Description,
		RecruiterId:  ju.RecruiterId,
		Requirements: convertedReqs,
	}
}
