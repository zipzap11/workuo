package routes

import (
	"workuo/factory"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	presenter := factory.Init()

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	// job
	e.POST("/jobs", presenter.JobPresentation.CreateJobPostHandler)
	e.GET("/jobs", presenter.JobPresentation.GetJobPostHandler)

	// user
	e.GET("/users", presenter.UserPresentation.GetUsersHandler)
	e.GET("/users/:id", presenter.UserPresentation.GetUserByIdHandler)
	e.POST("/users/register", presenter.UserPresentation.RegisterUserHandler)
	e.POST("/users/login", presenter.UserPresentation.LoginUserHandler)

	// recruiter
	e.POST("/recruiters/register", presenter.RecruiterPresentation.RegisterRecruiterHandler)
	e.POST("/recruiters/login", presenter.RecruiterPresentation.LoginRecruiterHandler)
	e.GET("/recruiters", presenter.RecruiterPresentation.GetRecruitersHandler)
	return e
}
