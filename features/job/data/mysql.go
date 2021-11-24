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
	err := jr.DB.Joins("JOIN recruiters ON jobs.recruiter_id = recruiters.id and company = ? and title like ?", data.Company, "%"+data.Title+"%").Preload("Requirements").Find(&jobs).Error
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

func (jr *mysqlJobRepository) DeleteJobData(data job.JobCore) error {
	err := jr.DB.Debug().Delete(&Job{}, data.ID).Error
	if err != nil {
		return err
	}

	err = jr.DB.Debug().Where("job_id = ?", data.ID).Delete(&Requirement{}).Error
	if err != nil {
		return err
	}

	return nil
}
