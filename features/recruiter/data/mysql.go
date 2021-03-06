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

func (rp *RecruiterRepository) CheckRecruiter(data recruiter.RecruiterCore) (recruiter.RecruiterCore, error) {
	var recruiterData Recruiter

	err := rp.DB.Where("email = ? and password = ?", data.Email, data.Password).First(&recruiterData).Error
	if err != nil {
		return recruiter.RecruiterCore{}, err
	}

	if recruiterData.ID == 0 && recruiterData.Email == "" {
		return recruiter.RecruiterCore{}, errors.New("no existing recruiter")
	}

	return ToCore(recruiterData), nil
}

func (rp *RecruiterRepository) GetRecruiters() ([]recruiter.RecruiterCore, error) {
	var recruiters []Recruiter

	err := rp.DB.Find(&recruiters).Error
	if err != nil {
		return nil, err
	}

	return ToCoreList(recruiters), nil
}

func (rp *RecruiterRepository) GetRecruiterById(data recruiter.RecruiterCore) (recruiter.RecruiterCore, error) {
	var recruiterData Recruiter

	err := rp.DB.First(&recruiterData, data.ID).Error
	if err != nil {
		return recruiter.RecruiterCore{}, err
	}

	return ToCore(recruiterData), nil
}
