package data

import (
	"workuo/features/job"

	"gorm.io/gorm"
)

type mysqlJobRepository struct {
	DB *gorm.DB
}

func NewMysqlJobRepository(DB *gorm.DB) job.Repository {
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
	err := jr.DB.Debug().Joins("JOIN recruiters ON jobs.recruiter_id = recruiters.id and company like ? and title like ?", "%"+data.Company+"%", "%"+data.Title+"%").Preload("Requirements").Find(&jobs).Error
	if err != nil {
		return nil, err
	}

	return toCoreList(jobs), nil
}

func (jr *mysqlJobRepository) GetJobDataById(id int) (job.JobCore, error) {
	var jobData Job
	err := jr.DB.Preload("Requirements").First(&jobData, id).Error

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

func (jr *mysqlJobRepository) UpdateJobData(data job.JobCore) error {
	jobData, requirements := SeparateJobRequirement(toRecordJob(data))

	err := jr.DB.Debug().Where("id = ?", data.ID).Updates(&jobData).Error
	if err != nil {
		return err
	}

	for _, req := range requirements {
		if req.ID != 0 {
			if req.Description == "" {
				err = jr.DB.Debug().Delete(&Requirement{}, req.ID).Error
				if err != nil {
					return err
				}
			} else {
				err = jr.DB.Debug().Model(&Requirement{}).Where("id = ?", req.ID).
					Update("description", req.Description).Error
				if err != nil {
					return err
				}
			}
		} else if req.ID == 0 {
			req.JobID = uint(data.ID)
			err = jr.DB.Debug().Select("JobID", "Description").Create(&req).Error
		}
		if err != nil {
			return err
		}
	}

	return nil
}
