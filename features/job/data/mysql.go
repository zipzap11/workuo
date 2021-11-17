package data

import (
	"workuo/features/job"
	"workuo/features/job/presentation/request"

	"gorm.io/gorm"
)

type mysqlJobRepository struct {
	DB *gorm.DB
}

func NewMysqlJobRepository(DB *gorm.DB) job.Data {
	return &mysqlJobRepository{DB}
}

func (j *mysqlJobRepository) InsertData(data request.Job) error {
	result := j.DB.Create(&data)
	if result.Error != nil {
		return result.Error
	}

	// return result.RowsAffected, result.Error
	return result.Error
}

func (j *mysqlJobRepository) SelectAllData() ([]job.JobCore, error) {
	var data []Job
	result := j.DB.Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return toCoreList(data), result.Error
}

func (j *mysqlJobRepository) UpdatedData(data job.JobCore) (job.JobCore, error) {
	var jobData Job
	j.DB.Debug().Where("id = ?", data.ID).First(&jobData)

	jobData.Description = data.Description
	jobData.Title = data.Title

	result := j.DB.Save(&jobData)
	if result.Error != nil {
		return job.JobCore{}, result.Error
	}

	return jobData.toCore(), nil
}

func (j *mysqlJobRepository) DeleteData(id int) error {
	var jobData Job

	err := j.DB.Delete(&jobData, id).Error
	if err != nil {
		return err
	}

	return nil
}
