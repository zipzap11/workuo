package service

import (
	"workuo/features/recruiter"
)

type recruiterService struct {
	recruiterRepository recruiter.Repository
}

func NewRecruiterService(recruiterRepo recruiter.Repository) recruiter.Service {
	return &recruiterService{recruiterRepo}
}

func (rs *recruiterService) RegisterRecruiter(data recruiter.RecruiterCore) error {
	err := rs.recruiterRepository.CreateRecruiter(data)
	if err != nil {
		return err
	}

	return nil
}

func (rs *recruiterService) LoginRecruiter(data recruiter.RecruiterCore) error {
	err := rs.recruiterRepository.CheckRecruiter(data)
	if err != nil {
		return err
	}

	return nil
}
