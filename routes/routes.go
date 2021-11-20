package routes

import (
	"workuo/factory"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	jobPresenter := factory.Init()
	userPresenter := factory.InitUser()
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.POST("/jobs", jobPresenter.JobPresentation.CreateJobPostHandler)
	e.POST("/users/register", userPresenter.UserHandler.RegisterUserHandler)
	e.GET("/users", userPresenter.UserHandler.GetAllUserHandler)
	e.POST("/users/login", userPresenter.UserHandler.LoginUserHandler)
	return e
}
