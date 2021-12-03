package service

import (
	"errors"
	"workuo/features/recruiter"
	"workuo/helper"
	"workuo/middleware"
)

type recruiterService struct {
	recruiterRepository recruiter.Repository
}

func NewRecruiterService(recruiterRepo recruiter.Repository) recruiter.Service {
	return &recruiterService{recruiterRepo}
}

func (rs *recruiterService) RegisterRecruiter(data recruiter.RecruiterCore) error {
	if !helper.ValidateEmail(data.Email) || !helper.ValidatePassword(data.Password) {
		return errors.New("incomplete or invalid data")
	}
	err := rs.recruiterRepository.CreateRecruiter(data)
	if err != nil {
		return err
	}

	return nil
}

func (rs *recruiterService) LoginRecruiter(data recruiter.RecruiterCore) (recruiter.RecruiterCore, error) {
	if !helper.ValidateEmail(data.Email) || !helper.ValidatePassword(data.Password) {
		return recruiter.RecruiterCore{}, errors.New("invalid data")
	}
	data, err := rs.recruiterRepository.CheckRecruiter(data)
	if err != nil {
		return recruiter.RecruiterCore{}, err
	}

	data.Token, err = middleware.CreateToken(data.ID, "recruiter")
	if err != nil {
		return recruiter.RecruiterCore{}, err
	}

	return data, nil
}

func (rs *recruiterService) GetRecruiters() ([]recruiter.RecruiterCore, error) {
	recruiters, err := rs.recruiterRepository.GetRecruiters()
	if err != nil {
		return nil, err
	}

	return recruiters, nil
}

func (rs *recruiterService) GetRecruiterById(data recruiter.RecruiterCore) (recruiter.RecruiterCore, error) {
	recruiterData, err := rs.recruiterRepository.GetRecruiterById(data)
	if err != nil {
		return recruiter.RecruiterCore{}, err
	}

	return recruiterData, nil
}
