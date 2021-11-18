package data

import (
	"workuo/features/job"

	"gorm.io/gorm"
)

type mysqlJobRepository struct {
	DB *gorm.DB
}

func NewMysqlJobRepository(DB *gorm.DB) job.Data {
	return &mysqlJobRepository{DB}
}

func (j *mysqlJobRepository) InsertData(data job.JobCore) error {
	// convertedRequirement := []Requirement{}
	// for _, req := range data.Requirements {
	// 	convertedRequirement = append(convertedRequirement, Requirement{
	// 		Description: req,
	// 	})
	// }
	// convertedData := Job{
	// 	Title:        data.Title,
	// 	Description:  data.Description,
	// 	Recruiter_id: data.RecruiterID,
	// 	Requirements: convertedRequirement,
	// }
	recordData := toRecordJob(data)
	result := j.DB.Create(&recordData)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
