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
	ID          uint `gorm: "primaryKey"`
	JobID       uint
	Description string
}

func toRecordRequirement(req job.RequirementCore) Requirement {
	return Requirement{
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
	return job.JobCore{
		ID:          int(j.ID),
		Title:       j.Title,
		Description: j.Description,
		RecruiterId: j.RecruiterId,
		Created_at:  j.CreatedAt,
		Updated_at:  j.UpdatedAt,
	}
}

func toCoreList(jobs []Job) []job.JobCore {
	var convertedData []job.JobCore
	for _, job := range jobs {
		convertedData = append(convertedData, job.toCore())
	}

	return convertedData
}
