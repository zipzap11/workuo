package service

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"workuo/features/job"
	"workuo/features/job/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	jobService    job.Service
	jobRepository mocks.Repository
	jobData       job.JobCore
	jobsData      []job.JobCore
)

func TestMain(m *testing.M) {
	jobService = NewJobUseCase(&jobRepository)
	jobData = job.JobCore{
		Title:       "Software Engineer",
		Description: "Create, develop, testing application",
		Company:     "Linkedin",
		Requirements: []job.RequirementCore{
			{
				Description: "Minimum 3 years experience",
			},
			{
				Description: "Strong knowledge on java and rust",
			},
		},
	}
	jobsData = []job.JobCore{
		{
			ID:          1,
			Title:       "Software Engineer",
			Description: "Create, develop, testing application",
			Company:     "Linkedin",
			Requirements: []job.RequirementCore{
				{
					Description: "Minimum 3 years experience",
				},
				{
					Description: "Strong knowledge on java and rust",
				},
			},
		},
	}

	os.Exit(m.Run())
}

func TestCreateJobPost(t *testing.T) {
	t.Run("create job success", func(t *testing.T) {
		jobRepository.On("InsertData", mock.AnythingOfType("job.JobCore")).Return(nil).Once()
		err := jobService.CreateJobPost(jobData)
		assert.Nil(t, err)
	})

	t.Run("create job failed invalid data", func(t *testing.T) {
		err := jobService.CreateJobPost(job.JobCore{})
		assert.NotNil(t, err)
		assert.Equal(t, "invalid data", err.Error())
	})

	t.Run("create job failed InsertData", func(t *testing.T) {
		jobRepository.On("InsertData", mock.AnythingOfType("job.JobCore")).Return(errors.New("error insert data")).Once()
		err := jobService.CreateJobPost(jobData)
		assert.NotNil(t, err)
		assert.Equal(t, "error insert data", err.Error())
	})
}

func TestGetJobPost(t *testing.T) {
	t.Run("get job post success", func(t *testing.T) {
		jobRepository.On("GetJobData", mock.AnythingOfType("job.JobCore")).Return(jobsData, nil).Once()
		resp, err := jobService.GetJobPost(jobData)
		assert.Nil(t, err)
		assert.Equal(t, len(jobsData), len(resp))
		assert.Equal(t, jobsData[0].Title, resp[0].Title)
	})

	t.Run("get job post error GetJobData", func(t *testing.T) {
		jobRepository.On("GetJobData", mock.AnythingOfType("job.JobCore")).Return(nil, errors.New("error get job data")).Once()
		resp, err := jobService.GetJobPost(jobData)
		assert.Nil(t, resp)
		assert.NotNil(t, err)
		assert.Equal(t, "error get job data", err.Error())
	})
}

func TestGetJobPostById(t *testing.T) {
	t.Run("get job post by id success", func(t *testing.T) {
		jobRepository.On("GetJobDataById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		resp, err := jobService.GetJobPostById(1)
		assert.Nil(t, err)
		assert.Equal(t, jobData.Title, resp.Title)
		assert.Equal(t, len(jobData.Requirements), len(resp.Requirements))
	})

	t.Run("get job post by id error GetJobDataById", func(t *testing.T) {
		jobRepository.On("GetJobDataById", mock.AnythingOfType("int")).Return(job.JobCore{}, errors.New("error get job data by id")).Once()
		resp, err := jobService.GetJobPostById(1)
		assert.NotNil(t, err)
		assert.Equal(t, "", resp.Title)
		assert.Equal(t, 0, len(resp.Requirements))
	})
}

func TestDeleteJobPost(t *testing.T) {
	t.Run("delete job post success", func(t *testing.T) {
		jobRepository.On("GetJobDataById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		jobRepository.On("DeleteJobData", mock.AnythingOfType("job.JobCore")).Return(nil).Once()
		err := jobService.DeleteJobPost(jobData)
		assert.Nil(t, err)
	})

	t.Run("delete job post error get job data by id", func(t *testing.T) {
		jobRepository.On("GetJobDataById", mock.AnythingOfType("int")).Return(job.JobCore{}, errors.New("error get job data by id")).Once()
		err := jobService.DeleteJobPost(jobData)
		assert.NotNil(t, err)
		assert.Equal(t, "error get job data by id", err.Error())
	})

	t.Run("delete job post error DeleteJobData", func(t *testing.T) {
		jobRepository.On("GetJobDataById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		jobRepository.On("DeleteJobData", mock.AnythingOfType("job.JobCore")).Return(errors.New("error delete job data")).Once()
		err := jobService.DeleteJobPost(jobData)
		assert.NotNil(t, err)
		assert.Equal(t, "error delete job data", err.Error())
	})

	t.Run("delete job post error recruiter id not equal", func(t *testing.T) {
		jobData.ID = 1
		jobData.RecruiterId = 1
		jobRepository.On("GetJobDataById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		jobRepository.On("DeleteJobData", mock.AnythingOfType("job.JobCore")).Return(errors.New("error delete job data")).Once()

		jobData.RecruiterId = 2
		err := jobService.DeleteJobPost(jobData)
		assert.NotNil(t, err)
		msg := fmt.Sprintf("recruiter with id %v does not have job with id %v", 2, 1)
		assert.Equal(t, msg, err.Error())
	})
}

func TestUpdateJobPost(t *testing.T) {
	t.Run("update job success", func(t *testing.T) {
		jobRepository.On("GetJobDataById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		jobRepository.On("UpdateJobData", mock.AnythingOfType("job.JobCore")).Return(nil).Once()
		err := jobService.UpdateJobPost(jobData)
		assert.Nil(t, err)
	})

	t.Run("update job error GetJobDataById", func(t *testing.T) {
		jobRepository.On("GetJobDataById", mock.AnythingOfType("int")).Return(job.JobCore{}, errors.New("error get data by id")).Once()
		err := jobService.UpdateJobPost(jobData)
		assert.NotNil(t, err)
		assert.Equal(t, "error get data by id", err.Error())
	})

	t.Run("update job error recruiter id not equal", func(t *testing.T) {
		jobData.ID = 1
		jobData.RecruiterId = 1
		jobRepository.On("GetJobDataById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		jobData.RecruiterId = 2
		err := jobService.UpdateJobPost(jobData)
		assert.NotNil(t, err)
		msg := fmt.Sprintf("recruiter with id %v does not have job with id %v", 2, 1)
		assert.Equal(t, msg, err.Error())
	})

	t.Run("update job error UpdateJobData", func(t *testing.T) {
		jobRepository.On("GetJobDataById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		jobRepository.On("UpdateJobData", mock.AnythingOfType("job.JobCore")).Return(errors.New("error update job data")).Once()
		err := jobService.UpdateJobPost(jobData)
		assert.NotNil(t, err)
		assert.Equal(t, "error update job data", err.Error())
	})

	t.Run("update job error invalid data", func(t *testing.T) {
		err := jobService.UpdateJobPost(job.JobCore{})
		assert.NotNil(t, err)
		assert.Equal(t, "invalid data", err.Error())
	})
}
