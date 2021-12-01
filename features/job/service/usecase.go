package service

import (
	"errors"
	"fmt"
	"workuo/features/job"
	"workuo/helper"
)

type jobUseCase struct {
	jobRepository job.Repository
}

func NewJobUseCase(jobRepository job.Repository) job.Service {
	return &jobUseCase{jobRepository}
}

func (ju *jobUseCase) CreateJobPost(data job.JobCore) error {
	if helper.IsEmpty(data.Title) || helper.IsEmpty(data.Description) {
		return errors.New("invalid data")
	}
	err := ju.jobRepository.InsertData(data)

	if err != nil {
		return err
	}

	return nil
}

func (ju *jobUseCase) GetJobPost(data job.JobCore) ([]job.JobCore, error) {
	jobData, err := ju.jobRepository.GetJobData(data)

	if err != nil {
		return nil, err
	}

	return jobData, nil
}

func (ju *jobUseCase) GetJobPostById(id int) (job.JobCore, error) {
	jobData, err := ju.jobRepository.GetJobDataById(id)

	if err != nil {
		return job.JobCore{}, err
	}

	return jobData, nil
}

func (ju *jobUseCase) DeleteJobPost(data job.JobCore) error {
	jobData, err := ju.jobRepository.GetJobDataById(data.ID)
	if err != nil {
		return err
	}
	if jobData.RecruiterId != data.RecruiterId {
		msg := fmt.Sprintf("recruiter with id %v does not have job with id %v", data.RecruiterId, data.ID)
		return errors.New(msg)
	}

	err = ju.jobRepository.DeleteJobData(data)
	if err != nil {
		return err
	}

	return nil
}

func (ju *jobUseCase) UpdateJobPost(data job.JobCore) error {
	if helper.IsEmpty(data.Title) || helper.IsEmpty(data.Description) {
		return errors.New("invalid data")
	}

	jobData, err := ju.jobRepository.GetJobDataById(data.ID)
	if err != nil {
		return err
	}
	if jobData.RecruiterId != data.RecruiterId {
		msg := fmt.Sprintf("recruiter with id %v does not have job with id %v", data.RecruiterId, data.ID)
		return errors.New(msg)
	}

	err = ju.jobRepository.UpdateJobData(data)
	if err != nil {
		return err
	}

	return nil
}
