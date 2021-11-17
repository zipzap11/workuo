package data

import (
	"workuo/features/job"

	"gorm.io/gorm"
)

type Job struct {
	// ID int `gorm: primaryKey`
	gorm.Model
	Title        string
	Description  string
	Recruiter_id int
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
