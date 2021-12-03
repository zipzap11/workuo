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

type ApplicationResponse struct {
	ID     uint
	UserID uint
	JobID  uint
	Status string
	Job    JobDetailResponse
	User   UserDetailResponse
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

type JobDetailResponse struct {
	ID           int
	Title        string
	Description  string
	RecruiterId  int
	Requirements []RequirementResponse
}

type RequirementResponse struct {
	ID          uint
	Description string
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

type UserDetailResponse struct {
	ID          uint
	Name        string
	Dob         time.Time
	Gender      string
	Address     string
	Title       string
	Bio         string
	Skillsets   []SkillsetResponse
	Experiences []ExperienceResponse
}

type SkillsetResponse struct {
	Name     string
	Category string
}

type ExperienceResponse struct {
	Title       string
	Description string
	StartDate   time.Time
	EndDate     time.Time
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

func ToApplicationResponse(data application.ApplicationCore) ApplicationResponse {
	return ApplicationResponse{
		ID:     data.ID,
		UserID: data.UserID,
		JobID:  data.JobID,
		Status: data.Status,
		Job:    ToJobDetailResponse(data.Job),
		User:   ToUserDetailResponse(data.User),
	}
}

func ToJobDetailResponse(data application.JobCore) JobDetailResponse {
	return JobDetailResponse{
		ID:           data.ID,
		Title:        data.Title,
		Description:  data.Description,
		RecruiterId:  data.RecruiterId,
		Requirements: ToRequirementsResponse(data.Requirements),
	}
}

func ToRequirementsResponse(data []application.RequirementCore) []RequirementResponse {
	converted := []RequirementResponse{}
	for _, req := range data {
		temp := RequirementResponse{req.ID, req.Description}
		converted = append(converted, temp)
	}
	return converted
}

func ToUserDetailResponse(data application.UserCore) UserDetailResponse {
	return UserDetailResponse{
		ID:          data.ID,
		Name:        data.Name,
		Dob:         data.Dob,
		Gender:      data.Gender,
		Address:     data.Address,
		Title:       data.Title,
		Bio:         data.Bio,
		Skillsets:   ToSkillsetResponse(data.Skillsets),
		Experiences: ToExperienceResponse(data.Experiences),
	}
}

func ToSkillsetResponse(data []application.SkillsetCore) []SkillsetResponse {
	convertedData := []SkillsetResponse{}
	for _, skill := range data {
		temp := SkillsetResponse{skill.Name, skill.Category}
		convertedData = append(convertedData, temp)
	}

	return convertedData
}

func ToExperienceResponse(data []application.ExperienceCore) []ExperienceResponse {
	convertedData := []ExperienceResponse{}
	for _, skill := range data {
		temp := ExperienceResponse{skill.Title, skill.Description, skill.StartDate, skill.EndDate}
		convertedData = append(convertedData, temp)
	}

	return convertedData
}
