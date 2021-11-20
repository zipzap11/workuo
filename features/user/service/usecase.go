package service

import (
	"fmt"
	"workuo/features/user"
)

type userService struct {
	userRepository user.Repository
}

func NewUserService(userRepository user.Repository) user.Service {
	return &userService{userRepository}
}

func (us *userService) RegisterUser(data user.UserCore) error {
	fmt.Println("data from service =====", data)
	err := us.userRepository.InsertData(data)

	if err != nil {
		return err
	}

	return nil
}

func (us *userService) GetAllUser() ([]user.UserCore, error) {
	users, err := us.userRepository.GetData()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (us *userService) LoginUser(data user.UserCore) (user.UserCore, error) {
	userData, err := us.userRepository.CheckUser(data)
	if err != nil {
		return user.UserCore{}, err
	}

	return userData, nil
}
