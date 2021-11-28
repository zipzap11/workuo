package service

import (
	"errors"
	"fmt"
	"time"
	"workuo/features/application"
	"workuo/features/job"
	"workuo/features/user"
)

type appService struct {
	appRepository application.Repository
	jobService    job.Service
	userService   user.Service
}

func NewAppService(ar application.Repository, js job.Service, us user.Service) application.Service {
	return &appService{
		appRepository: ar,
		jobService:    js,
		userService:   us,
	}
}

func (ar *appService) ApplyJob(data application.ApplicationCore) error {
	jobData, err := ar.jobService.GetJobPostById(int(data.JobID))
	if err != nil {
		return err
	}
	if jobData.ID == 0 {
		msg := fmt.Sprintf("job with id %v not found", data.JobID)
		return errors.New(msg)
	}

	appData, err := ar.appRepository.GetApplicationMultiParam(int(data.JobID), int(data.UserID))
	if err != nil {
		return err
	}
	if appData.ID != 0 {
		msg := fmt.Sprintf("user with id %v had applied job with id %v, current status = %v",
			appData.ID,
			appData.JobID,
			appData.Status,
		)
		return errors.New(msg)
	}

	data.Status = "pending"
	data.AppliedAt = time.Now()

	err = ar.appRepository.ApplyJob(data)
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

func (ar *appService) RejectApplication(id int, recruiterId int) error {
	data, err := ar.appRepository.GetApplicationByID(id)
	if err != nil {
		msg := fmt.Sprintf("application with id %v not found", id)
		return errors.New(msg)
	}
	if data.Job.RecruiterId != recruiterId {
		msg := fmt.Sprintf("recruiter with id %v not allowed to access post with id %v", recruiterId, id)
		return errors.New(msg)
	}
	if data.Status != "pending" {
		msg := fmt.Sprintf("this user has been %v", data.Status)
		return errors.New(msg)
	}

	err = ar.appRepository.RejectApplication(id)
	if err != nil {
		return err
	}

	return nil
}

func (ar *appService) AcceptApplication(id int, recruiterId int) error {
	data, err := ar.appRepository.GetApplicationByID(id)
	if err != nil {
		msg := fmt.Sprintf("application with id %v not found", id)
		return errors.New(msg)
	}
	if data.Job.RecruiterId != recruiterId {
		msg := fmt.Sprintf("recruiter with id %v not allowed to access post with id %v", recruiterId, id)
		return errors.New(msg)
	}
	if data.Status != "pending" {
		msg := fmt.Sprintf("this user has been %v", data.Status)
		return errors.New(msg)
	}

	err = ar.appRepository.AcceptApplication(id)
	if err != nil {
		return err
	}

	return nil
}

func (ar *appService) GetApplicationByID(id int) (application.ApplicationCore, error) {
	appData, err := ar.appRepository.GetApplicationByID(id)
	if err != nil {
		return application.ApplicationCore{}, err
	}

	userData, err := ar.userService.GetUserById(int(appData.UserID))
	if err != nil {
		return application.ApplicationCore{}, err
	}

	jobData, err := ar.jobService.GetJobPostById(int(appData.JobID))
	if err != nil {
		return application.ApplicationCore{}, err
	}

	appData.User = ToUserCore(userData)
	appData.Job = ToJobCore(jobData)

	return appData, nil
}

func (ar *appService) GetApplicationByJobID(id int) ([]application.ApplicationCore, error) {
	data, err := ar.appRepository.GetApplicationByJobID(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}
