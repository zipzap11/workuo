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

func (jr *mysqlJobRepository) InsertData(data job.JobCore) error {
	recordData := toRecordJob(data)
	result := jr.DB.Create(&recordData)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (jr *mysqlJobRepository) GetJobData(data job.JobCore) ([]job.JobCore, error) {
	var jobs []Job
	err := jr.DB.Preload("Requirements").Find(&jobs).Error
	if err != nil {
		return nil, err
	}

	return toCoreList(jobs), nil
}

func (jr *mysqlJobRepository) GetJobDataById(data job.JobCore) (job.JobCore, error) {
	var jobData Job
	err := jr.DB.Preload("Requirements").First(&jobData, data.ID).Error

	if err != nil {
		return job.JobCore{}, err
	}

	return jobData.toCore(), nil
}
