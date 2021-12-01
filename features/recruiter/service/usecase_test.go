package service

import (
	"errors"
	"os"
	"testing"
	"workuo/features/recruiter"
	"workuo/features/recruiter/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	recruiterRepo  mocks.Repository
	recruiterServ  recruiter.Service
	recruiterData  recruiter.RecruiterCore
	recruiterLogin recruiter.RecruiterCore
	recruitersData []recruiter.RecruiterCore
)

func TestMain(m *testing.M) {
	recruiterServ = NewRecruiterService(&recruiterRepo)
	recruiterData = recruiter.RecruiterCore{
		Company:  "Github",
		Address:  "New york city",
		Bio:      "Using the Hello World guide, you’ll create a repository, start a branch, write comments, and open a pull request.",
		Email:    "github@github.com",
		Password: "github",
	}
	recruiterLogin = recruiter.RecruiterCore{
		Email:    "github@github.com",
		Password: "github",
	}
	recruitersData = []recruiter.RecruiterCore{
		{
			ID:      1,
			Company: "Github",
			Address: "New york city",
			Email:   "github@github.com",
			Bio:     "Using the Hello World guide, you’ll create a repository, start a branch, write comments, and open a pull request.",
		},
	}
	os.Exit(m.Run())
}

func TestRegisterRecruiter(t *testing.T) {
	t.Run("Register recruiter success", func(t *testing.T) {
		recruiterRepo.On("CreateRecruiter", mock.AnythingOfType("recruiter.RecruiterCore")).Return(nil).Once()
		err := recruiterServ.RegisterRecruiter(recruiterData)
		assert.Nil(t, err)
	})

	t.Run("Register recruiter invalid data", func(t *testing.T) {
		err := recruiterServ.RegisterRecruiter(recruiter.RecruiterCore{})
		assert.NotNil(t, err)
		assert.Equal(t, "incomplete or invalid data", err.Error())
	})

	t.Run("Register recruiter invalid data", func(t *testing.T) {
		recruiterRepo.On("CreateRecruiter", mock.AnythingOfType("recruiter.RecruiterCore")).Return(errors.New("error create recruiter")).Once()
		err := recruiterServ.RegisterRecruiter(recruiterData)
		assert.NotNil(t, err)
		assert.Equal(t, "error create recruiter", err.Error())
	})

}

func TestLoginRecruiter(t *testing.T) {
	t.Run("Login recruiter success", func(t *testing.T) {
		recruiterRepo.On("CheckRecruiter", mock.AnythingOfType("recruiter.RecruiterCore")).Return(recruiterData, nil).Once()
		resp, err := recruiterServ.LoginRecruiter(recruiterLogin)
		assert.Nil(t, err)
		assert.Equal(t, recruiterData.Email, resp.Email)
		assert.Equal(t, recruiterData.Company, resp.Company)
	})

	t.Run("Login recruiter error invalid data", func(t *testing.T) {
		resp, err := recruiterServ.LoginRecruiter(recruiter.RecruiterCore{})
		assert.NotNil(t, err)
		assert.Equal(t, "invalid data", err.Error())
		assert.Equal(t, "", resp.Company)
		assert.Equal(t, "", resp.Email)
	})

	t.Run("Login recruiter error check recruiter", func(t *testing.T) {
		recruiterRepo.On("CheckRecruiter", mock.AnythingOfType("recruiter.RecruiterCore")).Return(recruiter.RecruiterCore{}, errors.New("error check recruiter")).Once()
		resp, err := recruiterServ.LoginRecruiter(recruiterLogin)
		assert.NotNil(t, err)
		assert.Equal(t, "error check recruiter", err.Error())
		assert.Equal(t, "", resp.Company)
		assert.Equal(t, "", resp.Email)
	})
}

func TestGetRecruiters(t *testing.T) {
	t.Run("get recruiters success", func(t *testing.T) {
		recruiterRepo.On("GetRecruiters").Return(recruitersData, nil).Once()
		resp, err := recruiterServ.GetRecruiters()
		assert.Nil(t, err)
		assert.Equal(t, len(recruitersData), len(resp))
		assert.Equal(t, recruitersData[0].Company, resp[0].Company)
	})

	t.Run("get recruiters error get recruiters", func(t *testing.T) {
		recruiterRepo.On("GetRecruiters").Return(nil, errors.New("error get recruiters")).Once()
		resp, err := recruiterServ.GetRecruiters()
		assert.NotNil(t, err)
		assert.Equal(t, "error get recruiters", err.Error())
		assert.Nil(t, resp)
	})
}

func TestGetRecruiterById(t *testing.T) {
	t.Run("get recruiter by id success", func(t *testing.T) {
		recruiterRepo.On("GetRecruiterById", mock.AnythingOfType("recruiter.RecruiterCore")).Return(recruiterData, nil).Once()
		resp, err := recruiterServ.GetRecruiterById(recruiter.RecruiterCore{ID: 1})
		assert.Nil(t, err)
		assert.Equal(t, recruiterData.Email, resp.Email)
		assert.Equal(t, recruiterData.Company, resp.Company)
	})

	t.Run("get recruiter by id error GetRecruiterById", func(t *testing.T) {
		recruiterRepo.On("GetRecruiterById", mock.AnythingOfType("recruiter.RecruiterCore")).Return(recruiter.RecruiterCore{}, errors.New("error get recruiter by id")).Once()
		resp, err := recruiterServ.GetRecruiterById(recruiter.RecruiterCore{ID: 1})
		assert.NotNil(t, err)
		assert.Equal(t, "error get recruiter by id", err.Error())
		assert.Equal(t, "", resp.Email)
		assert.Equal(t, "", resp.Company)
	})
}
