package service

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"time"
	"workuo/features/application"
	app_m "workuo/features/application/mocks"
	"workuo/features/invitation"
	"workuo/features/invitation/mocks"
	"workuo/features/job"
	job_m "workuo/features/job/mocks"
	"workuo/features/user"
	user_m "workuo/features/user/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	invService    invitation.Service
	invRepository mocks.Repository
	jobService    job_m.Service
	userService   user_m.Service
	appService    app_m.Service
	jobData       job.JobCore
	userData      user.UserCore
	appData       application.ApplicationCore
	invData       invitation.InvitationCore
	invsData      []invitation.InvitationCore
)

func TestMain(m *testing.M) {
	invService = NewInvitationService(&invRepository, &jobService, &userService, &appService)
	jobData = job.JobCore{
		ID:          1,
		RecruiterId: 1,
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
	userData = user.UserCore{
		Id:      1,
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
	appData = application.ApplicationCore{
		ID:        0,
		UserID:    1,
		JobID:     1,
		Status:    "pending",
		AppliedAt: time.Now(),
	}
	invData = invitation.InvitationCore{
		UserID:      1,
		JobID:       1,
		RecruiterID: 1,
		Role:        "recruiter",
		Status:      "pending",
	}
	invsData = []invitation.InvitationCore{
		invData,
	}
	os.Exit(m.Run())
}

func TestInviteUser(t *testing.T) {
	t.Run("invite user success", func(t *testing.T) {
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		appService.On("GetApplicationMultiParam", mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(appData, nil).Once()
		invRepository.On("InviteUser", mock.AnythingOfType("invitation.InvitationCore")).Return(nil).Once()
		err := invService.InviteUser(invData)
		assert.Nil(t, err)
	})

	t.Run("invite user error role not allowed", func(t *testing.T) {
		invData.Role = "user"
		err := invService.InviteUser(invData)
		assert.NotNil(t, err)
		assert.Equal(t, "only recruiter role allowed to invite user", err.Error())
	})

	t.Run("invite user error GetJobPostById", func(t *testing.T) {
		invData.Role = "recruiter"
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(job.JobCore{}, errors.New("error get job")).Once()
		err := invService.InviteUser(invData)
		assert.NotNil(t, err)
		assert.Equal(t, "error get job", err.Error())
	})

	t.Run("invite user error job not found", func(t *testing.T) {
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(job.JobCore{}, nil).Once()
		err := invService.InviteUser(invData)
		assert.NotNil(t, err)
		msg := fmt.Sprintf("job with id %v not found", 1)
		assert.Equal(t, msg, err.Error())
	})

	t.Run("invite user error rec id not equal", func(t *testing.T) {
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(job.JobCore{ID: 1, RecruiterId: 2}, nil).Once()
		err := invService.InviteUser(invData)
		assert.NotNil(t, err)
		msg := fmt.Sprintf("recruiter with id %v didn't have job post with id %v", 1, 1)
		assert.Equal(t, msg, err.Error())
	})

	t.Run("invite user error user id = 0", func(t *testing.T) {
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(user.UserCore{}, nil).Once()
		err := invService.InviteUser(invData)
		assert.NotNil(t, err)
		msg := fmt.Sprintf("user with id %v not found", 1)
		assert.Equal(t, msg, err.Error())
	})

	t.Run("invite user error GetUserByID", func(t *testing.T) {
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, errors.New("error get user")).Once()
		err := invService.InviteUser(invData)
		assert.NotNil(t, err)
		assert.Equal(t, "error get user", err.Error())
	})

	t.Run("invite user error has applied", func(t *testing.T) {
		appData.ID = 1
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		appService.On("GetApplicationMultiParam", mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(appData, nil).Once()
		err := invService.InviteUser(invData)
		assert.NotNil(t, err)
		msg := fmt.Sprintf("user with id %v has applied this job with status %v", 1, "pending")
		assert.Equal(t, msg, err.Error())
	})

	t.Run("invite user error has applied", func(t *testing.T) {
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		appService.On("GetApplicationMultiParam", mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(application.ApplicationCore{}, errors.New("error get application")).Once()
		err := invService.InviteUser(invData)
		assert.NotNil(t, err)
		assert.Equal(t, "error get application", err.Error())
	})

	t.Run("invite user error InviteUser", func(t *testing.T) {
		appData.ID = 0
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		appService.On("GetApplicationMultiParam", mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(appData, nil).Once()
		invRepository.On("InviteUser", mock.AnythingOfType("invitation.InvitationCore")).Return(errors.New("error invite user")).Once()
		err := invService.InviteUser(invData)
		assert.NotNil(t, err)
		assert.Equal(t, "error invite user", err.Error())
	})
}

func TestGetInvitationByID(t *testing.T) {
	t.Run("get invtiation by id success", func(t *testing.T) {
		invData.ID = 1
		invRepository.On("GetInvitationByID", mock.AnythingOfType("int")).Return(invData, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(jobData, nil).Once()

		resp, err := invService.GetInvitationByID(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(resp.ID))
		assert.Equal(t, "pending", resp.Status)
		assert.Equal(t, 1, int(resp.UserID))
	})

	t.Run("get invtiation by id error GetInvitatioByID", func(t *testing.T) {
		invData.ID = 1
		invRepository.On("GetInvitationByID", mock.AnythingOfType("int")).Return(invitation.InvitationCore{}, errors.New("error get invitation")).Once()
		resp, err := invService.GetInvitationByID(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, int(resp.ID))
		assert.Equal(t, "", resp.Status)
	})

	t.Run("get invtiation by id error inv not exist", func(t *testing.T) {
		invData.ID = 1
		invRepository.On("GetInvitationByID", mock.AnythingOfType("int")).Return(invitation.InvitationCore{}, nil).Once()
		resp, err := invService.GetInvitationByID(1)
		assert.NotNil(t, err)
		msg := fmt.Sprintf("invitation with id %v doesn't exist", 1)
		assert.Equal(t, msg, err.Error())
		assert.Equal(t, 0, int(resp.ID))
		assert.Equal(t, "", resp.Status)
	})

	t.Run("get invtiation by id error GetUserById", func(t *testing.T) {
		invData.ID = 1
		invRepository.On("GetInvitationByID", mock.AnythingOfType("int")).Return(invData, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(user.UserCore{}, errors.New("error get user")).Once()

		resp, err := invService.GetInvitationByID(1)
		assert.NotNil(t, err)
		assert.Equal(t, "error get user", err.Error())
		assert.Equal(t, 0, int(resp.ID))
		assert.Equal(t, "", resp.Status)
	})

	t.Run("get invtiation by id success", func(t *testing.T) {
		invData.ID = 1
		invRepository.On("GetInvitationByID", mock.AnythingOfType("int")).Return(invData, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(job.JobCore{}, errors.New("error get job")).Once()

		resp, err := invService.GetInvitationByID(1)
		assert.NotNil(t, err)
		assert.Equal(t, "error get job", err.Error())
		assert.Equal(t, 0, int(resp.ID))
		assert.Equal(t, "", resp.Status)
	})
}

func TestAcceptInvitation(t *testing.T) {
	invData.ID = 1
	t.Run("accept invitation success", func(t *testing.T) {
		invRepository.On("GetInvitationByID", mock.AnythingOfType("int")).Return(invData, nil).Once()
		invRepository.On("AcceptInvitation", mock.AnythingOfType("int")).Return(nil).Once()
		appService.On("ApplyJob", mock.AnythingOfType("application.ApplicationCore")).Return(nil).Once()
		err := invService.AcceptInvitation(1, 1)
		assert.Nil(t, err)
	})

	t.Run("accept invitation error get invitation", func(t *testing.T) {
		invRepository.On("GetInvitationByID", mock.AnythingOfType("int")).Return(invitation.InvitationCore{}, errors.New("error get invitation")).Once()
		err := invService.AcceptInvitation(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, "error get invitation", err.Error())
	})

	t.Run("accept invitation error invitation not found", func(t *testing.T) {
		invRepository.On("GetInvitationByID", mock.AnythingOfType("int")).Return(invitation.InvitationCore{}, nil).Once()
		err := invService.AcceptInvitation(1, 1)
		assert.NotNil(t, err)
		msg := fmt.Sprintf("invitation with id %v doesn't exist", 1)
		assert.Equal(t, msg, err.Error())
	})

	t.Run("accept invitation error user id incorrect", func(t *testing.T) {
		invData.UserID = 2
		invRepository.On("GetInvitationByID", mock.AnythingOfType("int")).Return(invData, nil).Once()
		err := invService.AcceptInvitation(1, 1)
		assert.NotNil(t, err)
		msg := fmt.Sprintf("user with id %v did not have invitation with id %v", 1, 1)
		assert.Equal(t, msg, err.Error())
	})

	t.Run("accept invitation error has already invited", func(t *testing.T) {
		invData.UserID = 1
		invData.Status = "accepted"
		invRepository.On("GetInvitationByID", mock.AnythingOfType("int")).Return(invData, nil).Once()
		err := invService.AcceptInvitation(1, 1)
		assert.NotNil(t, err)
		msg := fmt.Sprintf("user with id %v has %v invitation with id %v", 1, "accepted", 1)
		assert.Equal(t, msg, err.Error())
	})

	t.Run("accept invitation error acceptInvitation", func(t *testing.T) {
		invData.Status = "pending"
		invRepository.On("GetInvitationByID", mock.AnythingOfType("int")).Return(invData, nil).Once()
		invRepository.On("AcceptInvitation", mock.AnythingOfType("int")).Return(errors.New("error accept invitation")).Once()
		err := invService.AcceptInvitation(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, "error accept invitation", err.Error())
	})

	t.Run("accept invitation error apply job", func(t *testing.T) {
		invRepository.On("GetInvitationByID", mock.AnythingOfType("int")).Return(invData, nil).Once()
		invRepository.On("AcceptInvitation", mock.AnythingOfType("int")).Return(nil).Once()
		appService.On("ApplyJob", mock.AnythingOfType("application.ApplicationCore")).Return(errors.New("error apply job")).Once()
		err := invService.AcceptInvitation(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, "error apply job", err.Error())

	})

}

func TestRejectInvitation(t *testing.T) {
	t.Run("Reject invitation success", func(t *testing.T) {
		invData.ID = 1
		invRepository.On("GetInvitationByID", mock.AnythingOfType("int")).Return(invData, nil).Once()
		invRepository.On("RejectInvitation", mock.AnythingOfType("int")).Return(nil).Once()
		err := invService.RejectInvitation(1, 1)
		assert.Nil(t, err)
	})

	t.Run("Reject invitation success", func(t *testing.T) {
		invData.ID = 1
		invRepository.On("GetInvitationByID", mock.AnythingOfType("int")).Return(invitation.InvitationCore{}, errors.New("error get invitation")).Once()
		err := invService.RejectInvitation(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, "error get invitation", err.Error())
	})

	t.Run("Reject invitation error inv not found", func(t *testing.T) {
		invRepository.On("GetInvitationByID", mock.AnythingOfType("int")).Return(invitation.InvitationCore{}, nil).Once()
		err := invService.RejectInvitation(1, 1)
		assert.NotNil(t, err)
		msg := fmt.Sprintf("invitation with id %v doesn't exist", 1)
		assert.Equal(t, msg, err.Error())
	})

	t.Run("Reject invitation error user id not equals", func(t *testing.T) {
		invData.ID = 1
		invRepository.On("GetInvitationByID", mock.AnythingOfType("int")).Return(invData, nil).Once()
		err := invService.RejectInvitation(2, 1)
		assert.NotNil(t, err)
		msg := fmt.Sprintf("user with id %v did not have invitation with id %v", 2, 1)
		assert.Equal(t, msg, err.Error())
	})

	t.Run("Reject invitation error has accepted/rejected", func(t *testing.T) {
		invData.ID = 1
		invData.Status = "rejected"
		invRepository.On("GetInvitationByID", mock.AnythingOfType("int")).Return(invData, nil).Once()
		err := invService.RejectInvitation(1, 1)
		assert.NotNil(t, err)
		msg := fmt.Sprintf("user with id %v has %v invitation with id %v", 1, "rejected", 1)
		assert.Equal(t, msg, err.Error())
	})

	t.Run("Reject invitation error RejectInvitation", func(t *testing.T) {
		invData.ID = 1
		invData.Status = "pending"
		invRepository.On("GetInvitationByID", mock.AnythingOfType("int")).Return(invData, nil).Once()
		invRepository.On("RejectInvitation", mock.AnythingOfType("int")).Return(errors.New("error reject invitation")).Once()
		err := invService.RejectInvitation(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, "error reject invitation", err.Error())
	})
}

func TestGetInvitationByUserID(t *testing.T) {
	t.Run("Get invitation success", func(t *testing.T) {
		invRepository.On("GetInvitationByUserID", mock.AnythingOfType("int")).Return(invsData, nil).Once()
		resp, err := invService.GetInvitationByUserID(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(resp))
	})

	t.Run("Get invitation error", func(t *testing.T) {
		invRepository.On("GetInvitationByUserID", mock.AnythingOfType("int")).Return(nil, errors.New("error get invitation")).Once()
		resp, err := invService.GetInvitationByUserID(1)
		assert.NotNil(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, "error get invitation", err.Error())
	})
}

func TestGetInvitationByJobID(t *testing.T) {
	t.Run("get invitation by job id success", func(t *testing.T) {
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		invRepository.On("GetInvitationByJobID", mock.AnythingOfType("int")).Return(invsData, nil).Once()
		resp, err := invService.GetInvitationByJobID(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(resp))
		assert.Equal(t, invsData[0].Status, resp[0].Status)
	})

	t.Run("get invitation by job id error GetJobBypostId", func(t *testing.T) {
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(job.JobCore{}, errors.New("error get job")).Once()
		resp, err := invService.GetInvitationByJobID(1, 1)
		assert.NotNil(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, "error get job", err.Error())
		assert.Equal(t, 0, len(resp))
	})

	t.Run("get invitation by job id error recruiter id no equals", func(t *testing.T) {
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		resp, err := invService.GetInvitationByJobID(1, 2)
		assert.NotNil(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, 0, len(resp))
		msg := fmt.Sprintf("recruiter with id %v doesn't have access", 2)
		assert.Equal(t, msg, err.Error())
	})

	t.Run("get invitation by job id error GetInvitationByJobID", func(t *testing.T) {
		jobService.On("GetJobPostById", mock.AnythingOfType("int")).Return(jobData, nil).Once()
		invRepository.On("GetInvitationByJobID", mock.AnythingOfType("int")).Return(nil, errors.New("error get invitation")).Once()
		resp, err := invService.GetInvitationByJobID(1, 1)
		assert.NotNil(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, 0, len(resp))
		assert.Equal(t, "error get invitation", err.Error())
	})
}
