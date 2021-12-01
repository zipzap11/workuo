package service

import (
	"errors"
	"os"
	"testing"
	"time"
	"workuo/features/user"
	"workuo/features/user/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userRepo   mocks.Repository
	userServ   user.Service
	usersData  []user.UserCore
	userData   user.UserCore
	userLogin  user.UserCore
	userUpdate user.UserCore
)

func TestMain(m *testing.M) {
	userServ = NewUserService(&userRepo)

	usersData = []user.UserCore{
		{
			Id:       1,
			Name:     "Francisco",
			Dob:      time.Now(),
			Bio:      "test bio",
			Gender:   "Male",
			Address:  "Jakarta",
			Title:    "Software Engineer",
			Email:    "fran@gmail.com",
			Password: "fran123",
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
		},
	}

	userData = user.UserCore{
		Name:     "Francisco",
		Dob:      time.Now(),
		Bio:      "test bio",
		Gender:   "Male",
		Address:  "Jakarta",
		Title:    "Software Engineer",
		Email:    "fran@gmail.com",
		Password: "fran123",
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

	userLogin = user.UserCore{
		Email:    "fran@gmail.com",
		Password: "fran123",
	}

	userUpdate = user.UserCore{
		Name: "Rio",
		Experiences: []user.ExperienceCore{
			{
				Title:       "test",
				Description: "testing description",
				UserId:      1,
				StartDate:   time.Now(),
				EndDate:     time.Now().Add(time.Hour * 1),
			},
			{
				Id: 2,
			},
			{
				Id:          1,
				Title:       "test1",
				Description: "testing description1",
				StartDate:   time.Now(),
				EndDate:     time.Now().Add(time.Hour * 1),
			},
		},
		Skillsets: []user.SkillsetCore{
			{
				Name:     "test",
				Category: "test",
			},
			{
				Id: 2,
			},
			{
				Id:       1,
				Name:     "test1",
				Category: "test1",
			},
		},
	}
	os.Exit(m.Run())
}

func TestGetUser(t *testing.T) {
	t.Run("validate get users", func(t *testing.T) {
		userRepo.On("GetData", mock.AnythingOfType("user.UserCore")).Return(usersData, nil).Once()
		resp, err := userServ.GetUsers(user.UserCore{})
		assert.Nil(t, err)
		assert.Equal(t, len(resp), 1)
	})

	t.Run("error get users", func(t *testing.T) {
		userRepo.On("GetData", mock.AnythingOfType("user.UserCore")).Return(nil, errors.New("error on db"))
		resp, err := userServ.GetUsers(user.UserCore{})
		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})
}

func TestRegisterUser(t *testing.T) {
	t.Run("Register user success", func(t *testing.T) {
		userRepo.On("GetUserByEmail", mock.AnythingOfType("string")).Return(false, nil).Once()
		userRepo.On("InsertUserData", mock.AnythingOfType("user.UserCore")).Return(1, nil).Once()
		userRepo.On("CreateSkillset", mock.AnythingOfType("user.SkillsetCore")).Return(1, nil).Once()
		userRepo.On("CreateSkillset", mock.AnythingOfType("user.SkillsetCore")).Return(2, nil).Once()
		userRepo.On("AddUserSkillset", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil).Once()
		userRepo.On("AddUserSkillset", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil).Once()
		userRepo.On("CreateExperience", mock.AnythingOfType("user.ExperienceCore")).Return(nil).Once()

		err := userServ.RegisterUser(userData)
		assert.Nil(t, err)
	})

	t.Run("Register user error invalid email", func(t *testing.T) {
		// userRepo.On("GetUserByEmail", mock.AnythingOfType("string")).Return(false, errors.New("error")).Once()
		err := userServ.RegisterUser(user.UserCore{
			Email: "023ujawol",
		})
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "incomplete or invalid data")
	})

	t.Run("Register user error GetUserByEmail", func(t *testing.T) {
		userRepo.On("GetUserByEmail", mock.AnythingOfType("string")).Return(false, errors.New("error")).Once()
		err := userServ.RegisterUser(userData)
		assert.NotNil(t, err)
	})

	t.Run("Register user error email exist", func(t *testing.T) {
		userRepo.On("GetUserByEmail", mock.AnythingOfType("string")).Return(true, nil).Once()
		err := userServ.RegisterUser(userData)
		assert.NotNil(t, err)
	})

	t.Run("Register user error insert data", func(t *testing.T) {
		userRepo.On("GetUserByEmail", mock.AnythingOfType("string")).Return(false, nil).Once()
		userRepo.On("InsertUserData", mock.AnythingOfType("user.UserCore")).Return(0, errors.New("error")).Once()
		err := userServ.RegisterUser(userData)
		assert.NotNil(t, err)
	})

	t.Run("Register user error create skillset", func(t *testing.T) {
		userRepo.On("GetUserByEmail", mock.AnythingOfType("string")).Return(false, nil).Once()
		userRepo.On("InsertUserData", mock.AnythingOfType("user.UserCore")).Return(1, nil).Once()
		userRepo.On("CreateSkillset", mock.AnythingOfType("user.SkillsetCore")).Return(0, errors.New("error")).Once()
		err := userServ.RegisterUser(userData)
		assert.NotNil(t, err)
		assert.Equal(t, "error", err.Error())
	})

	t.Run("Register user error create user skillset", func(t *testing.T) {
		userRepo.On("GetUserByEmail", mock.AnythingOfType("string")).Return(false, nil).Once()
		userRepo.On("InsertUserData", mock.AnythingOfType("user.UserCore")).Return(1, nil).Once()
		userRepo.On("CreateSkillset", mock.AnythingOfType("user.SkillsetCore")).Return(1, nil).Once()
		userRepo.On("AddUserSkillset", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(errors.New("error add user skillset")).Once()
		err := userServ.RegisterUser(userData)
		assert.NotNil(t, err)
		assert.Equal(t, "error add user skillset", err.Error())
	})

	t.Run("Register user error create experiences", func(t *testing.T) {
		userRepo.On("GetUserByEmail", mock.AnythingOfType("string")).Return(false, nil).Once()
		userRepo.On("InsertUserData", mock.AnythingOfType("user.UserCore")).Return(1, nil).Once()
		userRepo.On("CreateSkillset", mock.AnythingOfType("user.SkillsetCore")).Return(1, nil).Once()
		userRepo.On("CreateSkillset", mock.AnythingOfType("user.SkillsetCore")).Return(2, nil).Once()
		userRepo.On("AddUserSkillset", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil).Once()
		userRepo.On("AddUserSkillset", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil).Once()
		userRepo.On("CreateExperience", mock.AnythingOfType("user.ExperienceCore")).Return(errors.New("error create experience")).Once()

		err := userServ.RegisterUser(userData)
		assert.NotNil(t, err)
		assert.Equal(t, "error create experience", err.Error())
	})
}

func TestLoginUser(t *testing.T) {
	t.Run("Login user success", func(t *testing.T) {
		userRepo.On("CheckUser", mock.AnythingOfType("user.UserCore")).Return(userData, nil).Once()
		data, err := userServ.LoginUser(userLogin)
		assert.Equal(t, userData.Email, data.Email)
		assert.Nil(t, err)
	})

	t.Run("Login failed email invalid", func(t *testing.T) {
		data, err := userServ.LoginUser(user.UserCore{
			Email:    "9012uhjja",
			Password: "fran123",
		})
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "invalid data")
		assert.Empty(t, data.Email)
	})

	t.Run("Login error check user", func(t *testing.T) {
		userRepo.On("CheckUser", mock.AnythingOfType("user.UserCore")).Return(user.UserCore{}, errors.New("error check data")).Once()
		data, err := userServ.LoginUser(userLogin)
		assert.Equal(t, "error check data", err.Error())
		assert.NotNil(t, err)
		assert.Empty(t, data.Id)
	})
}

func TestGetUserByID(t *testing.T) {
	t.Run("Get user by id success", func(t *testing.T) {
		userRepo.On("GetDataById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		data, err := userServ.GetUserById(1)
		assert.Equal(t, userData.Id, data.Id)
		assert.Nil(t, err)
	})

	t.Run("Get user by id error", func(t *testing.T) {
		userRepo.On("GetDataById", mock.AnythingOfType("int")).Return(user.UserCore{}, errors.New("error get user")).Once()
		data, err := userServ.GetUserById(1)
		assert.Empty(t, data)
		assert.NotNil(t, err)
		assert.Equal(t, "error get user", err.Error())
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("Update user success", func(t *testing.T) {
		userRepo.On("UpdateUser", mock.AnythingOfType("user.UserCore")).Return(nil).Once()
		userRepo.On("CreateSkillset", mock.AnythingOfType("user.SkillsetCore")).Return(2, nil).Once()
		userRepo.On("AddUserSkillset", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil).Once()
		userRepo.On("DeleteUserSkillset", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil).Once()
		userRepo.On("CreateSkillset", mock.AnythingOfType("user.SkillsetCore")).Return(1, nil).Once()
		userRepo.On("UpdateUserSkillset", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil).Once()
		userRepo.On("CreateExperience", mock.AnythingOfType("user.ExperienceCore")).Return(nil).Once()
		userRepo.On("DeleteExperience", mock.AnythingOfType("int")).Return(nil).Once()
		userRepo.On("UpdateExperience", mock.AnythingOfType("user.ExperienceCore")).Return(nil).Once()

		err := userServ.UpdateUser(userUpdate)
		assert.Nil(t, err)
	})

	t.Run("Update user error UpdateUserData", func(t *testing.T) {
		userRepo.On("UpdateUser", mock.AnythingOfType("user.UserCore")).Return(errors.New("error update user")).Once()
		err := userServ.UpdateUser(userUpdate)
		assert.NotNil(t, err)
		assert.Equal(t, "error update user", err.Error())
	})

	t.Run("Update user success", func(t *testing.T) {
		userRepo.On("UpdateUser", mock.AnythingOfType("user.UserCore")).Return(nil).Once()
		userRepo.On("CreateSkillset", mock.AnythingOfType("user.SkillsetCore")).Return(0, errors.New("error create skillset")).Once()
		err := userServ.UpdateUser(userUpdate)
		assert.NotNil(t, err)
		assert.Equal(t, "error create skillset", err.Error())
	})
}
