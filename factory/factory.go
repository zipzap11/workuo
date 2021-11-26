package factory

import (
	"workuo/driver"
	//job domain
	jobData "workuo/features/job/data"
	jobPresent "workuo/features/job/presentation"
	jobService "workuo/features/job/service"

	// user domain
	userData "workuo/features/user/data"
	userPresent "workuo/features/user/presentation"
	userService "workuo/features/user/service"

	// recruiter domain
	recruiterData "workuo/features/recruiter/data"
	recruiterPresent "workuo/features/recruiter/presentation"
	recruiterService "workuo/features/recruiter/service"

	// recruiter domain
	invitationData "workuo/features/invitation/data"
	invitationPresent "workuo/features/invitation/presentation"
	invitationService "workuo/features/invitation/service"
)

type jobPresenter struct {
	JobPresentation        jobPresent.JobHandler
	UserPresentation       userPresent.UserHandler
	RecruiterPresentation  recruiterPresent.RecruiterHandler
	InvitationPresentation invitationPresent.InvitationHandler
}

func Init() jobPresenter {
	// job layer
	jobData := jobData.NewMysqlJobRepository(driver.DB)
	jobService := jobService.NewJobUseCase(jobData)

	// user layer
	userData := userData.NewMysqlUserRepository(driver.DB)
	userService := userService.NewUserService(userData)

	// recruiter layer
	recruiterData := recruiterData.NewRecruiterRepository(driver.DB)
	recruiterService := recruiterService.NewRecruiterService(recruiterData)

	// invitation layer
	invData := invitationData.NewInvitationRepository(driver.DB)
	invService := invitationService.NewInvitationService(invData)

	return jobPresenter{
		JobPresentation:        *jobPresent.NewJobHandler(jobService),
		UserPresentation:       *userPresent.NewUserHandler(userService),
		RecruiterPresentation:  *recruiterPresent.NewRecruiterHandler(recruiterService),
		InvitationPresentation: *invitationPresent.NewInvitationHandler(invService),
	}
}
