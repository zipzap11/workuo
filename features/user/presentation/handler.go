package presentation

import (
	"net/http"
	"strconv"
	"workuo/features/user"
	"workuo/features/user/presentation/request"
	"workuo/features/user/presentation/response"
	"workuo/middleware"

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

func (uh *UserHandler) GetUsersHandler(e echo.Context) error {
	var filter request.UserFilter
	err := e.Bind(&filter)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	data, err := uh.userService.GetUsers(user.UserCore{
		Title:     filter.Title,
		Skillsets: request.ToSkillsetsCore(filter.Skillsets),
	})
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

	data, err := uh.userService.GetUserById(id)
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

func (uh *UserHandler) UpdateUserHandler(e echo.Context) error {
	var userData request.UserRequest
	err := e.Bind(&userData)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	claims := middleware.ExtractClaim(e)
	userId := int(claims["id"].(float64))
	role := claims["role"]
	if role != "user" {
		return e.JSON(http.StatusForbidden, map[string]interface{}{
			"message": "role not allowed to update user data",
		})
	}

	userData.ID = userId

	err = uh.userService.UpdateUser(userData.ToUserCore())
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})

}
