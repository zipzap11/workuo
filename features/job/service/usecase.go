package service

import (
	// "errors"
	"workuo/features/job"
	"workuo/features/job/presentation/request"
)

type jobUseCase struct {
	jobRepository job.Data
}

func NewJobUseCase(jobRepository job.Data) job.Service {
	return &jobUseCase{jobRepository}
}

func (ju *jobUseCase) CreateJobPost(data request.Job) error {
	err := ju.jobRepository.InsertData(data)

	if err != nil {
		return err
	}

	return nil
}
