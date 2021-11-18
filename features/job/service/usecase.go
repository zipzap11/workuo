package service

import (
	// "errors"
	"fmt"
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
	fmt.Println("data from service", data)
	err := ju.jobRepository.InsertData(data)

	if err != nil {
		fmt.Println("error insert data")
		return err
	}

	return nil
}

// func (ju *jobUseCase) UpdateJobPost(data job.JobCore) (job.JobCore, error) {
// 	if data.Title == "" || data.Description == "" {
// 		return job.JobCore{}, errors.New("invalid data")
// 	}
// 	result, err := ju.jobRepository.UpdatedData(data)

// 	if err != nil {
// 		return job.JobCore{}, err
// 	}

// 	return result, nil
// }

// func (ju *jobUseCase) DeleteJobPost(id int) error {
// 	err := ju.jobRepository.DeleteData(id)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (ju *jobUseCase) GetAllJobPost() ([]job.JobCore, error) {
// 	data, err := ju.jobRepository.SelectAllData()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return data, nil
// }
