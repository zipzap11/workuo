package service

import (
	"workuo/features/recruiter"
	"workuo/middleware"
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

func (rs *recruiterService) LoginRecruiter(data recruiter.RecruiterCore) (recruiter.RecruiterCore, error) {
	data, err := rs.recruiterRepository.CheckRecruiter(data)
	if err != nil {
		return recruiter.RecruiterCore{}, err
	}

	data.Token, err = middleware.CreateToken(data.ID, data.Company)
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
