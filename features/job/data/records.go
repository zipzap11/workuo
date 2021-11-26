package data

import (
	"workuo/features/job"

	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Title        string
	Description  string
	RecruiterId  int
	Requirements []Requirement
}

type Requirement struct {
	gorm.Model
	JobID       uint
	Description string
}

func toRecordRequirement(req job.RequirementCore) Requirement {
	return Requirement{
		ID:          req.ID,
		JobID:       req.JobId,
		Description: req.Description,
	}
}

func toRecordJob(data job.JobCore) Job {
	convertedRequirement := []Requirement{}
	for _, req := range data.Requirements {
		convertedRequirement = append(convertedRequirement, toRecordRequirement(req))
	}
	return Job{
		Title:        data.Title,
		Description:  data.Description,
		RecruiterId:  data.RecruiterId,
		Requirements: convertedRequirement,
	}
}

func (j *Job) toCore() job.JobCore {
	convertedRequirement := []job.RequirementCore{}
	for _, req := range j.Requirements {
		convertedRequirement = append(convertedRequirement, req.toCore())
	}
	return job.JobCore{
		ID:           int(j.ID),
		Title:        j.Title,
		Description:  j.Description,
		RecruiterId:  j.RecruiterId,
		Requirements: convertedRequirement,
		Created_at:   j.CreatedAt,
		Updated_at:   j.UpdatedAt,
	}
}

func (r *Requirement) toCore() job.RequirementCore {
	return job.RequirementCore{
		ID:          r.ID,
		JobId:       r.JobID,
		Description: r.Description,
	}
}

func toCoreList(jobs []Job) []job.JobCore {
	var convertedData []job.JobCore
	for _, job := range jobs {
		convertedData = append(convertedData, job.toCore())
	}

	return convertedData
}

func SeparateJobRequirement(data Job) (Job, []Requirement) {
	newJob := Job{
		Title:       data.Title,
		Description: data.Description,
		RecruiterId: data.RecruiterId,
	}
	newRequirements := data.Requirements
	return newJob, newRequirements
}
