package service

import (
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

func (ar *appService) RejectApplication(id int) error {
	err := ar.appRepository.RejectApplication(id)
  if err != nil {
    return err
  }

	return nil
}

func (ar *appService) AcceptApplication(id int) error {
	err := ar.appRepository.AcceptApplication(id)
	if err != nil {
		return err
	}

	return nil
}
