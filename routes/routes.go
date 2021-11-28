package routes

import (
	"workuo/config"
	"workuo/factory"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	presenter := factory.Init()

	e := echo.New()
	jwt := e.Group("")
	jwt.Use(middleware.JWT([]byte(config.JWT_KEY)))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// job
	jwt.POST("/jobs", presenter.JobPresentation.CreateJobPostHandler)
	jwt.GET("/jobs", presenter.JobPresentation.GetJobPostHandler)
	jwt.GET("/jobs/:id", presenter.JobPresentation.GetJobPostByIdHandler)
	jwt.DELETE("/jobs/:id", presenter.JobPresentation.DeleteJobPostHandler)
	jwt.PUT("/jobs", presenter.JobPresentation.UpdateJobPostHandler)

	// user
	jwt.GET("/users", presenter.UserPresentation.GetUsersHandler)
	jwt.GET("/users/:id", presenter.UserPresentation.GetUserByIdHandler)
	e.POST("/users/register", presenter.UserPresentation.RegisterUserHandler)
	e.POST("/users/login", presenter.UserPresentation.LoginUserHandler)

	// recruiter
	e.POST("/recruiters/register", presenter.RecruiterPresentation.RegisterRecruiterHandler)
	e.POST("/recruiters/login", presenter.RecruiterPresentation.LoginRecruiterHandler)
	jwt.GET("/recruiters", presenter.RecruiterPresentation.GetRecruitersHandler)
	jwt.GET("/recruiters/:id", presenter.RecruiterPresentation.GetRecruiterByIdHandler)

	// invitation
	jwt.POST("/invitations", presenter.InvitationPresentation.InviteUserHandler)
	jwt.GET("/invitations/:id", presenter.InvitationPresentation.GetInvitationByIDHandler)
	jwt.GET("/invitations/users", presenter.InvitationPresentation.GetInvitationByUserID)
	jwt.GET("/invitations/jobs/:id", presenter.InvitationPresentation.GetInvitationByJobID)
	jwt.PUT("/invitations/accept", presenter.InvitationPresentation.AcceptInvitationHandler)
	jwt.PUT("/invitations/reject", presenter.InvitationPresentation.RejectInvitationHandler)

	// application
	jwt.POST("/applications", presenter.ApplicationPresentation.ApplyJobHandler)
	jwt.GET("/applications/:id", presenter.ApplicationPresentation.GetApplicationByIDHandler)
	jwt.GET("/applications/users/:id", presenter.ApplicationPresentation.GetApplicationByUserIdHandler)
	jwt.GET("/applications/jobs/:id", presenter.ApplicationPresentation.GetApplicationByJobIDHandler)
	jwt.PUT("/applications/reject", presenter.ApplicationPresentation.RejectApplicationHandler)
	jwt.PUT("/applications/accept", presenter.ApplicationPresentation.AcceptApplication)

	return e
}
