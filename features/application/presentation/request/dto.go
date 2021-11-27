package request

import "workuo/features/application"

type ApplicationRequest struct {
	JobID  uint `json: "JobID`
	UserID uint `json: "UserID`
}

func (app *ApplicationRequest) ToCore() application.ApplicationCore {
	return application.ApplicationCore{
		UserID: app.UserID,
		JobID:  app.JobID,
	}
}
