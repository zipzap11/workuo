package factory

import (
	"workuo/driver"
	"workuo/features/job/data"
	"workuo/features/job/presentation"
	"workuo/features/job/service"
)

type jobPresenter struct {
	jobPresentation presentation.JobHandler
}

func Init() jobPresenter {
	DB := driver.InitDB()
	jobData := data.NewMysqlJobRepository(DB)
	jobService := service.NewJobUseCase(jobData)

	return jobPresenter{jobPresentation: *presentation.NewJobHandler(jobService)}
}
