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

func (ar *appService) GetApplicationByUserID(id int) ([]application.ApplicationCore, error) {
	applications, err := ar.appRepository.GetApplicationByUserID(id)
	if err != nil {
		return nil, err
	}

	return applications, nil
}
