package data

import (
	"errors"
	"workuo/features/application"

	"gorm.io/gorm"
)

type mysqlAppRepository struct {
	DB *gorm.DB
}

func NewMysqlAppRepository(db *gorm.DB) application.Repository {
	return &mysqlAppRepository{db}
}

func (ar *mysqlAppRepository) ApplyJob(data application.ApplicationCore) error {
	appData := ToApplicationRecord(data)

	err := ar.DB.Create(&appData).Error
	if err != nil {
		return err
	}

	return nil
}

func (ar *mysqlAppRepository) GetApplicationByUserID(id int) ([]application.ApplicationCore, error) {
	var applications []Application
	err := ar.DB.Debug().Where("user_id = ?", id).Joins("Job").Preload("Job.Requirements").Find(&applications).Error
	if err != nil {
		return nil, err
	}

	return ToCoreList(applications), nil
}

func (ar *mysqlAppRepository) RejectApplication(id int) error {
	err := ar.DB.Model(&Application{}).Where("id = ?", id).Update("status", "rejected").Error
	if err != nil {
		return err
	}

	return nil
}

func (ar *mysqlAppRepository) AcceptApplication(id int) error {
	err := ar.DB.Model(&Application{}).Where("id = ?", id).Update("status", "accepted").Error
	if err != nil {
		return err
	}

	return nil
}

func (ar *mysqlAppRepository) GetApplicationByID(id int) (application.ApplicationCore, error) {
	var data Application
	err := ar.DB.Debug().First(&data, id).Error
	if err != nil {
		return application.ApplicationCore{}, err
	}
	if data.ID == 0 {
		return application.ApplicationCore{}, errors.New("application doesn't exist")
	}

	return ToCore(data), nil
}

func (ar *mysqlAppRepository) GetApplicationByJobID(id int) ([]application.ApplicationCore, error) {
	var applications []Application
	err := ar.DB.Where("job_id = ?", id).Joins("User").Find(&applications).Error
	if err != nil {
		return nil, err
	}

	return ToCoreList(applications), nil
}

func (ar *mysqlAppRepository) GetApplicationMultiParam(jobId int, userId int) (application.ApplicationCore, error) {
	var data Application
	err := ar.DB.Where("job_id = ? AND user_id = ?", jobId, userId).Find(&data).Error
	if err != nil {
		return application.ApplicationCore{}, err
	}

	return ToCore(data), nil
}
