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
