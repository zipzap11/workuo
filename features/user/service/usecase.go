package service

import "workuo/features/user"

type userService struct {
	userRepository user.Repository
}

func NewUserService(userRepository user.Repository) user.Service {
	return &userService{}
}

func (us *userService) RegisterUser(data user.UserCore) error {
	err := us.userRepository.InsertData(data)

	if err != nil {
		return err
	}

	return nil
}
