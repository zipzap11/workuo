package routes

import (
	"fmt"
	"net/http"
	"workuo/driver"
	"workuo/factory"
	"workuo/features/user/data"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	presenter := factory.Init()

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.POST("/jobs", presenter.JobPresentation.CreateJobPostHandler)
	e.POST("/users/register", registerUserController)
	return e
}

func registerUserController(e echo.Context) error {
	data := data.User{}
	err := e.Bind(&data)
	fmt.Println("data ======", data)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	err = driver.DB.Create(&data).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}
