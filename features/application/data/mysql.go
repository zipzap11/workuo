package data

import (
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
