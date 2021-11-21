package data

import (
	"errors"
	"workuo/features/recruiter"

	"gorm.io/gorm"
)

type RecruiterRepository struct {
	DB *gorm.DB
}

func NewRecruiterRepository(DB *gorm.DB) recruiter.Repository {
	return &RecruiterRepository{DB}
}

func (rp *RecruiterRepository) CreateRecruiter(data recruiter.RecruiterCore) error {
	convertedData := FromCore(data)
	err := rp.DB.Create(&convertedData).Error
	if err != nil {
		return err
	}

	return nil
}

func (rp *RecruiterRepository) CheckRecruiter(data recruiter.RecruiterCore) error {
	var recruiterData Recruiter

	err := rp.DB.Where("email = ? and password = ?", data.Email, data.Password).First(&recruiterData).Error
	if err != nil {
		return err
	}

	if recruiterData.ID == 0 && recruiterData.Email == "" {
		return errors.New("no existing recruiter")
	}

	return nil
}
