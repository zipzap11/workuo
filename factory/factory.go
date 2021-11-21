package factory

import (
	"workuo/driver"
	jobData "workuo/features/job/data"
	jobPresent "workuo/features/job/presentation"
	jobService "workuo/features/job/service"
	userData "workuo/features/user/data"
	userPresent "workuo/features/user/presentation"
	userService "workuo/features/user/service"
)

type jobPresenter struct {
	JobPresentation  jobPresent.JobHandler
	UserPresentation userPresent.UserHandler
}

func Init() jobPresenter {
	// job layer
	jobData := jobData.NewMysqlJobRepository(driver.DB)
	jobService := jobService.NewJobUseCase(jobData)

	// user layer
	userData := userData.NewMysqlUserRepository(driver.DB)
	userService := userService.NewUserService(userData)

	return jobPresenter{
		JobPresentation:  *jobPresent.NewJobHandler(jobService),
		UserPresentation: *userPresent.NewUserHandler(userService),
	}
}
