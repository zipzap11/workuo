package data

import (
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