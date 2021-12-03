package factory

import (
	"log"
	"workuo/config"
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

	// application domain
	applicationData "workuo/features/application/data"
	applicationPresent "workuo/features/application/presentation"
	applicationService "workuo/features/application/service"

	// invitation domain
	invitationData "workuo/features/invitation/data"
	invitationPresent "workuo/features/invitation/presentation"
	invitationService "workuo/features/invitation/service"

	// invitation domain
	newsData "workuo/features/news/data"
	newsPresent "workuo/features/news/presentation"
	newsService "workuo/features/news/service"
)

type jobPresenter struct {
	JobPresentation         jobPresent.JobHandler
	UserPresentation        userPresent.UserHandler
	RecruiterPresentation   recruiterPresent.RecruiterHandler
	ApplicationPresentation applicationPresent.AppHandler
	InvitationPresentation  invitationPresent.InvitationHandler
	NewsPresentation        newsPresent.NewsHandler
}

func Init() jobPresenter {
	configAPP, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	// job layer
	jobData := jobData.NewMysqlJobRepository(driver.DB)
	jobService := jobService.NewJobUseCase(jobData)

	// user layer
	userData := userData.NewMysqlUserRepository(driver.DB)
	userService := userService.NewUserService(userData)

	// recruiter layer
	recruiterData := recruiterData.NewRecruiterRepository(driver.DB)
	recruiterService := recruiterService.NewRecruiterService(recruiterData)

	// application layer
	appData := applicationData.NewMysqlAppRepository(driver.DB)
	appService := applicationService.NewAppService(appData, jobService, userService)

	// invitation layer
	invData := invitationData.NewInvitationRepository(driver.DB)
	invService := invitationService.NewInvitationService(invData, jobService, userService, appService)

	// news 3rd party api layer
	newsData := newsData.NewNewsApiRepository("http://api.mediastack.com/v1/news", configAPP.NewsAPIKey)
	newsService := newsService.NewApiService(newsData)

	return jobPresenter{
		JobPresentation:         *jobPresent.NewJobHandler(jobService),
		UserPresentation:        *userPresent.NewUserHandler(userService),
		RecruiterPresentation:   *recruiterPresent.NewRecruiterHandler(recruiterService),
		ApplicationPresentation: *applicationPresent.NewAppHandler(appService),
		InvitationPresentation:  *invitationPresent.NewInvitationHandler(invService),
		NewsPresentation:        *newsPresent.NewNewsHandler(newsService),
	}
}
