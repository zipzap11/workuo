package presentation

import (
	"fmt"
	"net/http"
	"strconv"
	"workuo/features/user"
	"workuo/features/user/presentation/request"
	"workuo/features/user/presentation/response"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *UserHandler {
	return &UserHandler{userService}
}

func (uh *UserHandler) RegisterUserHandler(e echo.Context) error {
	userData := request.UserRequest{}

	err := e.Bind(&userData)
	fmt.Println("data in handler =====", userData)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	err = uh.userService.RegisterUser(userData.ToUserCore())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

func (uh *UserHandler) GetAllUserHandler(e echo.Context) error {
	data, err := uh.userService.GetAllUser()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToUserResponseList(data),
	})

}

func (uh *UserHandler) LoginUserHandler(e echo.Context) error {
	userAuth := request.UserAuth{}
	e.Bind(&userAuth)
	data, err := uh.userService.LoginUser(userAuth.ToUserCore())

	if err != nil {
		return e.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToUserLoginResponse(data),
	})

}

func (uh *UserHandler) GetUserByIdHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	data, err := uh.userService.GetUserById(user.UserCore{Id: uint(id)})
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToUserResponse(data),
	})

}

func (uh *UserHandler) GetUserByTitleHandler(e echo.Context) error {
	title := e.QueryParam("title")

	data, err := uh.userService.GetUserByTitle(user.UserCore{Title: title})
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToUserResponseList(data),
	})

}
