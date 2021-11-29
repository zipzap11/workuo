package service

import (
	"workuo/features/user"
	"workuo/middleware"
)

type userService struct {
	userRepository user.Repository
}

func NewUserService(userRepository user.Repository) user.Service {
	return &userService{userRepository}
}

func (us *userService) RegisterUser(data user.UserCore) error {
	err := us.userRepository.InsertData(data)

	if err != nil {
		return err
	}

	return nil
}

func (us *userService) GetUsers(data user.UserCore) ([]user.UserCore, error) {
	users, err := us.userRepository.GetData(data)
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

	userData.Token, err = middleware.CreateToken(userData.Id, "user")
	if err != nil {
		return user.UserCore{}, err
	}

	return userData, nil
}

func (us *userService) GetUserById(id int) (user.UserCore, error) {
	userData, err := us.userRepository.GetDataById(id)

	if err != nil {
		return user.UserCore{}, err
	}

	return userData, nil
}

func (us *userService) UpdateUser(data user.UserCore) error {
	// update primary user data
	err := us.userRepository.UpdateUser(data)
	if err != nil {
		return err
	}

	// add or update or delete user-skillsets
	for _, skill := range data.Skillsets {
		if skill.Id == 0 {
			skillId, err := us.userRepository.CreateSkillset(skill)
			if err != nil {
				return err
			}

			err = us.userRepository.AddUserSkillset(int(data.Id), skillId)
			if err != nil {
				return err
			}
		} else {
			if skill.Name == "" {
				err := us.userRepository.DeleteUserSkillset(int(data.Id), int(skill.Id))
				if err != nil {
					return err
				}
			} else {
				skillId, err := us.userRepository.CreateSkillset(skill)
				if err != nil {
					return err
				}
				err = us.userRepository.UpdateUserSkillset(int(data.Id), int(skill.Id), skillId)
				if err != nil {
					return err
				}
			}
		}
	}

	// update experiences
	for _, exp := range data.Experiences {
		if exp.Id == 0 {
			exp.UserId = data.Id
			err = us.userRepository.CreateExperience(exp)
			if err != nil {
				return err
			}
		} else {
			if exp.Title == "" {
				err = us.userRepository.DeleteExperience(int(exp.Id))
			} else {
				err = us.userRepository.UpdateExperience(exp)
			}
			if err != nil {
				return err
			}
		}
	}

	return nil
}
