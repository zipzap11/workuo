package response

import (
	"net/http"
	"time"
	"workuo/features/application"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string
	Data    interface{}
}

func NewErrorResponse(e echo.Context, msg string, code int) error {
	return e.JSON(code, Response{
		Message: msg,
	})
}

func NewSuccessResponse(e echo.Context, msg string, data interface{}) error {
	return e.JSON(http.StatusOK, Response{
		Message: msg,
		Data:    data,
	})
}

type ApplicationResponseUser struct {
	ID     uint
	UserID uint
	JobID  uint
	Status string
	Job    JobResponse
}

type ApplicationResponseJob struct {
	ID     uint
	UserID uint
	JobID  uint
	Status string
	User   UserResponse
}

type JobResponse struct {
	ID          uint   `json: "id"`
	Title       string `json: "title"`
	Description string `json: "description`
}

type UserResponse struct {
	ID      uint
	Name    string
	Dob     time.Time
	Gender  string
	Address string
	Title   string
	Bio     string
}

func ToApplicationResponseUser(data application.ApplicationCore) ApplicationResponseUser {
	return ApplicationResponseUser{
		ID:     data.ID,
		UserID: data.UserID,
		JobID:  data.JobID,
		Status: data.Status,
		Job:    ToJobResponse(data.Job),
	}
}

func ToJobResponse(data application.JobCore) JobResponse {
	return JobResponse{
		ID:          uint(data.ID),
		Title:       data.Title,
		Description: data.Description,
	}
}

func ToUserResponse(data application.UserCore) UserResponse {
	return UserResponse{
		ID:      data.ID,
		Name:    data.Name,
		Dob:     data.Dob,
		Gender:  data.Gender,
		Address: data.Address,
		Title:   data.Title,
		Bio:     data.Bio,
	}
}

func ToApplicationResponseJob(data application.ApplicationCore) ApplicationResponseJob {
	return ApplicationResponseJob{
		ID:     data.ID,
		UserID: data.UserID,
		JobID:  data.JobID,
		Status: data.Status,
		User:   ToUserResponse(data.User),
	}
}

func ToApplicationResponseJobList(data []application.ApplicationCore) []ApplicationResponseJob {
	convertedData := []ApplicationResponseJob{}
	for _, app := range data {
		convertedData = append(convertedData, ToApplicationResponseJob(app))
	}

	return convertedData
}

func ToApplicationResponseUserList(data []application.ApplicationCore) []ApplicationResponseUser {
	convertedData := []ApplicationResponseUser{}
	for _, app := range data {
		convertedData = append(convertedData, ToApplicationResponseUser(app))
	}

	return convertedData
}
