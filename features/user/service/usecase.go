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
