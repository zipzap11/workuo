package service

import (
	"errors"
	"fmt"
	"time"
	"workuo/features/application"
)

type appService struct {
	appRepository application.Repository
}

func NewAppService(ar application.Repository) application.Service {
	return &appService{ar}
}

func (ar *appService) ApplyJob(data application.ApplicationCore) error {
	data.Status = "pending"
	data.AppliedAt = time.Now()
	err := ar.appRepository.ApplyJob(data)
	if err != nil {
		return err
	}

	return nil
}

func (ar *appService) GetApplicationByUserID(id int) ([]application.ApplicationCore, error) {
	applications, err := ar.appRepository.GetApplicationByUserID(id)
	if err != nil {
		return nil, err
	}

	return applications, nil
}

func (ar *appService) RejectApplication(id int) error {
	err := ar.appRepository.RejectApplication(id)
	if err != nil {
		return err
	}

	return nil
}

func (ar *appService) AcceptApplication(id int, recruiterId int) error {
	data, err := ar.appRepository.GetApplicationByID(id)
	if err != nil {
		return err
	}
	if data.Job.RecruiterId != recruiterId {
		msg := fmt.Sprintf("recruiter with id %v not allowed to access post with id %v", recruiterId, id)
		return errors.New(msg)
	}

	err = ar.appRepository.AcceptApplication(id)
	if err != nil {
		return err
	}

	return nil
}

func (ar *appService) GetApplicationByID(id int) (application.ApplicationCore, error) {
	data, err := ar.appRepository.GetApplicationByID(id)
	if err != nil {
		return application.ApplicationCore{}, err
	}

	return data, nil
}

func (ar *appService) GetApplicationByJobID(id int) ([]application.ApplicationCore, error) {
	data, err := ar.appRepository.GetApplicationByJobID(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}
