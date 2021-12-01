package service

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"time"
	"workuo/features/application"
	app_m "workuo/features/application/mocks"
	"workuo/features/job"
	job_m "workuo/features/job/mocks"
	"workuo/features/user"
	user_m "workuo/features/user/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	appServ       application.Service
	appRepository app_m.Repository
	userService   user_m.Service
	appMockServ   app_m.Service
	jobService    job_m.Service
	jobData       job.JobCore
	appData       application.ApplicationCore
	userData      user.UserCore
	appsData      []application.ApplicationCore
)

func TestMain(m *testing.M) {
	appServ = NewAppService(&appRepository, &jobService, &userService)
	jobData = job.JobCore{
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
	}
	appData = application.ApplicationCore{
		ID:        1,
		UserID:    1,
		JobID:     1,
		Status:    "pending",
		AppliedAt: time.Now(),
		Job: application.JobCore{
			Title:       "Software Engineer",
			Description: "Create, develop, testing application",
			Company:     "Linkedin",
			Requirements: []application.RequirementCore{
				{
					Description: "Minimum 3 years experience",
				},
				{
					Description: "Strong knowledge on java and rust",
				},
			},
		},
	}
	userData = user.UserCore{
		Name:    "Francisco",
		Dob:     time.Now(),
		Bio:     "test bio",
		Gender:  "Male",
		Address: "Jakarta",
		Title:   "Software Engineer",
		Skillsets: []user.SkillsetCore{
			{
				Name:     "HTML",
				Category: "WEB",
			},
			{
				Name:     "go",
				Category: "server",
			},
		},
		Experiences: []user.ExperienceCore{
			{
				UserId:      1,
				Title:       "Internship Program",
				Description: "intership",
				StartDate:   time.Now(),
				EndDate:     time.Now().Add(time.Hour * 1),
			},
		},
	}
	appsData = []application.ApplicationCore{
		appData,
	}
	os.Exit(m.Run())
}

func TestApplyJob(t *testing.T) {
	t.Run("apply job success", func(t *testing.T) {
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		appRepository.On("GetApplicationMultiParam", mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(application.ApplicationCore{
				ID: 0,
			}, nil).Once()
		appRepository.On("ApplyJob", mock.AnythingOfType("application.ApplicationCore")).Return(nil).Once()
		err := appServ.ApplyJob(application.ApplicationCore{
			UserID: 1,
			JobID:  1,
		})

		assert.Nil(t, err)
	})

	t.Run("apply job error GetJobPostById", func(t *testing.T) {
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(job.JobCore{}, errors.New("error get job by id")).Once()
		err := appServ.ApplyJob(application.ApplicationCore{
			UserID: 1,
			JobID:  1,
		})

		assert.NotNil(t, err)
		assert.Equal(t, "error get job by id", err.Error())
	})

	t.Run("apply job error jobdata id = 0", func(t *testing.T) {
		jobData.ID = 0
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		err := appServ.ApplyJob(application.ApplicationCore{
			UserID: 1,
			JobID:  1,
		})

		assert.NotNil(t, err)
		msg := fmt.Sprintf("job with id %v not found", 1)
		assert.Equal(t, msg, err.Error())
	})

	t.Run("apply job error GetApplicationMultiParam", func(t *testing.T) {
		jobData.ID = 1
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		appRepository.On("GetApplicationMultiParam", mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(application.ApplicationCore{}, errors.New("error get application")).Once()
		err := appServ.ApplyJob(application.ApplicationCore{
			UserID:    1,
			JobID:     1,
			AppliedAt: time.Now(),
		})

		assert.NotNil(t, err)
		assert.Equal(t, "error get application", err.Error())
	})

	t.Run("apply job error same application exist", func(t *testing.T) {
		jobData.ID = 1
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		appRepository.On("GetApplicationMultiParam", mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(application.ApplicationCore{
				ID:     1,
				UserID: 1,
				JobID:  1,
				Status: "pending",
			}, nil).Once()
		err := appServ.ApplyJob(application.ApplicationCore{
			UserID: 1,
			JobID:  1,
		})
		assert.NotNil(t, err)
		msg := fmt.Sprintf("user with id %v had applied job with id %v, current status = %v", 1, 1, "pending")
		assert.Equal(t, msg, err.Error())
	})

	t.Run("apply job error ApplyJob", func(t *testing.T) {
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		appRepository.On("GetApplicationMultiParam", mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(application.ApplicationCore{
				ID: 0,
			}, nil).Once()
		appRepository.On("ApplyJob", mock.AnythingOfType("application.ApplicationCore")).Return(errors.New("error apply job")).Once()
		err := appServ.ApplyJob(application.ApplicationCore{
			UserID: 1,
			JobID:  1,
		})

		assert.NotNil(t, err)
		assert.Equal(t, "error apply job", err.Error())
	})
}

func TestGetApplicationByUserID(t *testing.T) {
	applications := []application.ApplicationCore{
		{
			ID:        1,
			UserID:    1,
			JobID:     1,
			Status:    "pending",
			AppliedAt: time.Now(),
			Job: application.JobCore{
				ID:          1,
				Title:       "Software Engineer",
				Description: "test",
			},
		},
	}
	t.Run("get application by user id success", func(t *testing.T) {
		appRepository.On("GetApplicationByUserID", mock.AnythingOfType("int")).Return(applications, nil).Once()
		resp, err := appServ.GetApplicationByUserID(1)
		assert.Nil(t, err)
		assert.Equal(t, len(applications), len(resp))
		assert.Equal(t, applications[0].ID, resp[0].ID)
	})

	t.Run("get application by user id error GetApplicationByUserID", func(t *testing.T) {
		appRepository.On("GetApplicationByUserID", mock.AnythingOfType("int")).Return(nil, errors.New("error get application")).Once()
		resp, err := appServ.GetApplicationByUserID(1)
		assert.Nil(t, resp)
		assert.NotNil(t, err)
		assert.Equal(t, "error get application", err.Error())
	})
}

func TestRejectApplication(t *testing.T) {
	t.Run("reject application success", func(t *testing.T) {
		jobData.RecruiterId = 1
		appMockServ.On("GetApplicationByID", mock.AnythingOfType("int")).
			Return(application.ApplicationCore{
				ID:     1,
				UserID: 1,
				Status: "pending",
				Job: application.JobCore{
					ID:          1,
					RecruiterId: 1,
				},
			}, nil).Once()
		appRepository.On("GetApplicationByID", mock.AnythingOfType("int")).
			Return(application.ApplicationCore{
				ID:     1,
				UserID: 1,
				Status: "pending",
				Job: application.JobCore{
					ID:          1,
					RecruiterId: 1,
				},
			}, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		appRepository.On("RejectApplication", mock.AnythingOfType("int")).Return(nil).Once()

		err := appServ.RejectApplication(1, 1)
		assert.Nil(t, err)
	})

	t.Run("reject application error GetApplicationByID", func(t *testing.T) {
		appMockServ.On("GetApplicationByID", mock.AnythingOfType("int")).
			Return(application.ApplicationCore{}, errors.New("error")).Once()
		appRepository.On("GetApplicationByID", mock.AnythingOfType("int")).
			Return(application.ApplicationCore{}, errors.New("error")).Once()
		err := appServ.RejectApplication(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, "application with id 1 not found", err.Error())
	})

	t.Run("reject application error no match recruiter id", func(t *testing.T) {
		appMockServ.On("GetApplicationByID", mock.AnythingOfType("int")).
			Return(application.ApplicationCore{
				ID:     1,
				UserID: 1,
				Status: "pending",
				Job: application.JobCore{
					ID:          1,
					RecruiterId: 2,
				},
			}, nil).Once()
		appRepository.On("GetApplicationByID", mock.AnythingOfType("int")).
			Return(application.ApplicationCore{
				ID:     1,
				UserID: 1,
				Status: "pending",
				Job: application.JobCore{
					ID:          1,
					RecruiterId: 1,
				},
			}, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		err := appServ.RejectApplication(1, 2)
		msg := fmt.Sprintf("recruiter with id %v not allowed to access post with id %v", 2, 1)
		assert.NotNil(t, err)
		assert.Equal(t, msg, err.Error())
	})

	t.Run("reject application error already rejected or accepted", func(t *testing.T) {
		appMockServ.On("GetApplicationByID", mock.AnythingOfType("int")).
			Return(application.ApplicationCore{
				ID:     1,
				UserID: 1,
				Status: "accepted",
				Job: application.JobCore{
					ID:          1,
					RecruiterId: 1,
				},
			}, nil).Once()
		appRepository.On("GetApplicationByID", mock.AnythingOfType("int")).
			Return(application.ApplicationCore{
				ID:     1,
				UserID: 1,
				Status: "rejected",
				Job: application.JobCore{
					ID:          1,
					RecruiterId: 1,
				},
			}, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		err := appServ.RejectApplication(1, 1)
		msg := fmt.Sprintf("this user has been %v", "rejected")
		assert.NotNil(t, err)
		assert.Equal(t, msg, err.Error())
	})

	t.Run("reject application error RejectApplication", func(t *testing.T) {
		appMockServ.On("GetApplicationByID", mock.AnythingOfType("int")).Return(application.ApplicationCore{
			ID:     1,
			UserID: 1,
			JobID:  1,
			Status: "pending",
			Job: application.JobCore{
				ID:          1,
				RecruiterId: 1,
			},
		}, nil).Once()
		appRepository.On("GetApplicationByID", mock.AnythingOfType("int")).
			Return(application.ApplicationCore{
				ID:     1,
				UserID: 1,
				Status: "rejected",
				Job: application.JobCore{
					ID:          1,
					RecruiterId: 1,
				},
			}, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		appRepository.On("RejectApplication", mock.AnythingOfType("int")).Return(errors.New("error reject application")).Once()
		err := appServ.RejectApplication(1, 1)
		assert.NotNil(t, err)
	})

}

func TestAcceptApplication(t *testing.T) {
	t.Run("accept application success", func(t *testing.T) {
		jobData.RecruiterId = 1
		appMockServ.On("GetApplicationByID", mock.AnythingOfType("int")).Return(application.ApplicationCore{
			ID:     1,
			UserID: 1,
			JobID:  1,
			Job: application.JobCore{
				RecruiterId: 1,
			},
			Status: "pending",
		}, nil).Once()
		appRepository.On("GetApplicationByID", mock.AnythingOfType("int")).
			Return(application.ApplicationCore{
				ID:     1,
				UserID: 1,
				Status: "pending",
				Job: application.JobCore{
					ID:          1,
					RecruiterId: 1,
				},
			}, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		appRepository.On("AcceptApplication", mock.AnythingOfType("int")).Return(nil).Once()

		err := appServ.AcceptApplication(1, 1)
		assert.Nil(t, err)
	})

	t.Run("accept application error GetAPplicationByID", func(t *testing.T) {
		appMockServ.On("GetApplicationByID", mock.AnythingOfType("int")).Return(application.ApplicationCore{}, errors.New("error get appliaction")).Once()
		appRepository.On("GetApplicationByID", mock.AnythingOfType("int")).
			Return(application.ApplicationCore{}, errors.New("error")).Once()
		err := appServ.AcceptApplication(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, "application with id 1 not found", err.Error())
	})

	t.Run("accept application error recruiter id not match", func(t *testing.T) {
		appMockServ.On("GetApplicationByID", mock.AnythingOfType("int")).Return(application.ApplicationCore{
			ID:     1,
			UserID: 1,
			JobID:  1,
			Job: application.JobCore{
				RecruiterId: 1,
			},
			Status: "pending",
		}, nil).Once()
		appRepository.On("GetApplicationByID", mock.AnythingOfType("int")).
			Return(application.ApplicationCore{
				ID:     1,
				UserID: 1,
				Status: "pending",
				Job: application.JobCore{
					ID:          1,
					RecruiterId: 1,
				},
			}, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		err := appServ.AcceptApplication(1, 2)
		assert.NotNil(t, err)
		msg := fmt.Sprintf("recruiter with id %v not allowed to access application with id %v", 2, 1)
		assert.Equal(t, msg, err.Error())
	})

	t.Run("accept application error application accepted/rejected", func(t *testing.T) {
		appMockServ.On("GetApplicationByID", mock.AnythingOfType("int")).Return(application.ApplicationCore{
			ID:     1,
			UserID: 1,
			JobID:  1,
			Job: application.JobCore{
				RecruiterId: 1,
			},
			Status: "accepted",
		}, nil).Once()
		appRepository.On("GetApplicationByID", mock.AnythingOfType("int")).
			Return(application.ApplicationCore{
				ID:     1,
				UserID: 1,
				Status: "accepted",
				Job: application.JobCore{
					ID:          1,
					RecruiterId: 1,
				},
			}, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		err := appServ.AcceptApplication(1, 1)
		assert.NotNil(t, err)
		msg := fmt.Sprintf("this user has been %v", "accepted")
		assert.Equal(t, msg, err.Error())
	})

	t.Run("accept application error acceptApplication", func(t *testing.T) {
		appMockServ.On("GetApplicationByID", mock.AnythingOfType("int")).Return(application.ApplicationCore{
			ID:     1,
			UserID: 1,
			JobID:  1,
			Job: application.JobCore{
				RecruiterId: 1,
			},
			Status: "pending",
		}, nil).Once()
		appRepository.On("GetApplicationByID", mock.AnythingOfType("int")).
			Return(application.ApplicationCore{
				ID:     1,
				UserID: 1,
				Status: "accepted",
				Job: application.JobCore{
					ID:          1,
					RecruiterId: 1,
				},
			}, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		appRepository.On("AcceptApplication", mock.AnythingOfType("int")).Return(errors.New("error accept application")).Once()

		err := appServ.AcceptApplication(1, 1)
		assert.NotNil(t, err)
	})
}

func TestGetApplicationByID(t *testing.T) {
	t.Run("get application by id ", func(t *testing.T) {
		appRepository.On("GetApplicationByID", mock.AnythingOfType("int")).Return(appData, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		resp, err := appServ.GetApplicationByID(1)
		assert.Nil(t, err)
		assert.Equal(t, appData.ID, resp.ID)
		assert.Equal(t, 1, resp.Job.RecruiterId)
		assert.Equal(t, appData.User.ID, resp.User.ID)
	})

	t.Run("get application by id error GetApplicationByID", func(t *testing.T) {
		appRepository.On("GetApplicationByID", mock.AnythingOfType("int")).Return(application.ApplicationCore{}, errors.New("error get appliaction")).Once()
		resp, err := appServ.GetApplicationByID(1)
		assert.NotNil(t, err)
		assert.Equal(t, "error get appliaction", err.Error())
		assert.Equal(t, "", resp.Status)
	})

	t.Run("get application by id error GetUserById", func(t *testing.T) {
		appRepository.On("GetApplicationByID", mock.AnythingOfType("int")).Return(appData, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(user.UserCore{}, errors.New("error get user")).Once()
		resp, err := appServ.GetApplicationByID(1)
		assert.NotNil(t, err)
		assert.Equal(t, "error get user", err.Error())
		assert.Equal(t, "", resp.Status)
	})

	t.Run("get application by id error GetUserById", func(t *testing.T) {
		appRepository.On("GetApplicationByID", mock.AnythingOfType("int")).Return(appData, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(job.JobCore{}, errors.New("error get job")).Once()
		resp, err := appServ.GetApplicationByID(1)
		assert.NotNil(t, err)
		assert.Equal(t, "error get job", err.Error())
		assert.Equal(t, "", resp.Status)
	})

}

func TestGetApplicationByJobID(t *testing.T) {
	t.Run("get application by job id", func(t *testing.T) {
		appRepository.On("GetApplicationByJobID", mock.AnythingOfType("int")).Return(appsData, nil).Once()
		resp, err := appServ.GetApplicationByJobID(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(resp))
		assert.Equal(t, 1, int(resp[0].ID))
		assert.Equal(t, "pending", resp[0].Status)
	})

	t.Run("get application by job id error", func(t *testing.T) {
		appRepository.On("GetApplicationByJobID", mock.AnythingOfType("int")).Return(nil, errors.New("error get application")).Once()
		resp, err := appServ.GetApplicationByJobID(1)
		assert.NotNil(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, 0, len(resp))
	})
}

func TestGetApplicationMultiParam(t *testing.T) {
	t.Run("get application with multi param", func(t *testing.T) {
		appRepository.On("GetApplicationMultiParam", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(appData, nil).Once()
		resp, err := appServ.GetApplicationMultiParam(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(resp.ID))
		assert.Equal(t, "pending", resp.Status)
	})

	t.Run("get application with multiparam error", func(t *testing.T) {
		appRepository.On("GetApplicationMultiParam", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(application.ApplicationCore{}, errors.New("error get application")).Once()
		resp, err := appServ.GetApplicationMultiParam(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, int(resp.ID))
		assert.Equal(t, "", resp.Status)
	})
}
