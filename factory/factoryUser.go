package factory

import (
	"workuo/driver"
	"workuo/features/user/data"
	"workuo/features/user/presentation"
	"workuo/features/user/service"
)

type userPresenter struct {
	UserHandler presentation.UserHandler
}

func InitUser() userPresenter {
	userRepository := data.NewMysqlUserRepository(driver.DB)
	userService := service.NewUserService(userRepository)

	return userPresenter{
		UserHandler: *presentation.NewUserHandler(userService),
	}
}
