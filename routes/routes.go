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
	e.POST("/jobs", presenter.JobPresentation.CreateJobPostHandler)
	e.GET("/jobs", presenter.JobPresentation.GetJobPostHandler)
	e.GET("/jobs/:id", presenter.JobPresentation.GetJobPostByIdHandler)
	e.DELETE("/jobs/:id", presenter.JobPresentation.DeleteJobPostHandler)
	e.PUT("/jobs", presenter.JobPresentation.UpdateJobPostHandler)

	// user
	e.GET("/users", presenter.UserPresentation.GetUsersHandler)
	e.GET("/users/:id", presenter.UserPresentation.GetUserByIdHandler)
	e.POST("/users/register", presenter.UserPresentation.RegisterUserHandler)
	e.POST("/users/login", presenter.UserPresentation.LoginUserHandler)

	// recruiter
	e.POST("/recruiters/register", presenter.RecruiterPresentation.RegisterRecruiterHandler)
	e.POST("/recruiters/login", presenter.RecruiterPresentation.LoginRecruiterHandler)
	e.GET("/recruiters", presenter.RecruiterPresentation.GetRecruitersHandler)
	e.GET("/recruiters/:id", presenter.RecruiterPresentation.GetRecruiterByIdHandler)

	// invitation
	jwt.POST("/invitations", presenter.InvitationPresentation.InviteUserHandler)

	// application
	e.POST("/applications", presenter.ApplicationPresentation.ApplyJobHandler)
	e.GET("/applications/:id", presenter.ApplicationPresentation.GetApplicationByIDHandler)
	e.GET("/applications/users/:id", presenter.ApplicationPresentation.GetApplicationByUserIdHandler)
	e.GET("/applications/jobs/:id", presenter.ApplicationPresentation.GetApplicationByJobIDHandler)
	e.PUT("/applications/reject", presenter.ApplicationPresentation.RejectApplicationHandler)
	e.PUT("/applications/accept", presenter.ApplicationPresentation.AcceptApplication)

	return e
}
