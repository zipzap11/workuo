package data

import (
	"workuo/features/job"
	"workuo/features/job/presentation/request"

	"gorm.io/gorm"
)

type Job struct {
	// ID int `gorm: primaryKey`
	gorm.Model
	Title        string
	Description  string
	Recruiter_id int
	Requirements []Requirement
}

type Requirement struct {
	ID          uint `gorm: "primaryKey"`
	JobID       uint
	Description string
}

func toRecord(requestData request.Job) Job {
	convertedRequirement := []Requirement{}
	for _, req := range requestData.Requirements {
		convertedRequirement = append(convertedRequirement, Requirement{
			Description: req,
		})
	}
	return Job{
		Title:        requestData.Title,
		Description:  requestData.Description,
		Recruiter_id: requestData.RecruiterID,
		Requirements: convertedRequirement,
	}
}

func (j *Job) toCore() job.JobCore {
	return job.JobCore{
		ID:           int(j.ID),
		Title:        j.Title,
		Description:  j.Description,
		Recruiter_id: j.Recruiter_id,
		Created_at:   j.CreatedAt,
		Updated_at:   j.UpdatedAt,
	}
}

func toCoreList(jobs []Job) []job.JobCore {
	var convertedData []job.JobCore
	for _, job := range jobs {
		convertedData = append(convertedData, job.toCore())
	}

	return convertedData
}
