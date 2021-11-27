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
