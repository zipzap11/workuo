package service

import (
	"workuo/features/job"
)

type jobUseCase struct {
	jobRepository job.Data
}

func NewJobUseCase(jobRepository job.Data) job.Service {
	return &jobUseCase{jobRepository}
}

func (ju *jobUseCase) CreateJobPost(data job.JobCore) error {
	err := ju.jobRepository.InsertData(data)

	if err != nil {
		return err
	}

	return nil
}
