package response

import (
	"net/http"
	"time"
	"workuo/features/invitation"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string      `json: "message"`
	Data    interface{} `json: "data"`
}

func NewSuccessResponse(e echo.Context, data interface{}) error {
	return e.JSON(http.StatusOK, Response{
		Message: "Success",
		Data:    data,
	})
}

func NewErrorResponse(e echo.Context, code int, msg string) error {
	return e.JSON(code, Response{
		Message: msg,
	})
}

type InvitationDetailResponse struct {
	ID          uint
	RecruiterID uint
	Status      string
	User        UserDetailResponse
	Job         JobDetailResponse
}

func ToInvitationDetailResponse(data invitation.InvitationCore) InvitationDetailResponse {
	return InvitationDetailResponse{
		ID:          data.ID,
		RecruiterID: data.ID,
		Status:      data.Status,
		User:        ToUserDetailResponse(data.User),
		Job:         ToJobResponseDetail(data.Job),
	}
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
	Description string
	Title       string
	StartDate   time.Time
	EndDate     time.Time
}

func ToUserDetailResponse(data invitation.UserCore) UserDetailResponse {
	return UserDetailResponse{
		ID:          data.ID,
		Name:        data.Name,
		Dob:         data.Dob,
		Gender:      data.Gender,
		Address:     data.Address,
		Title:       data.Title,
		Bio:         data.Bio,
		Skillsets:   ToSkillsetsResponse(data.Skillsets),
		Experiences: ToExperiencesResponse(data.Experiences),
	}
}

func ToSkillsetsResponse(data []invitation.SkillsetCore) []SkillsetResponse {
	converted := []SkillsetResponse{}
	for _, data := range data {
		temp := SkillsetResponse{
			Name:     data.Name,
			Category: data.Category,
		}
		converted = append(converted, temp)
	}
	return converted
}

func ToExperiencesResponse(data []invitation.ExperienceCore) []ExperienceResponse {
	converted := []ExperienceResponse{}
	for _, ex := range data {
		temp := ExperienceResponse{
			Title:       ex.Title,
			Description: ex.Description,
			StartDate:   ex.StartDate,
			EndDate:     ex.EndDate,
		}
		converted = append(converted, temp)
	}

	return converted
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
	JobId       uint
	Description string
}

func ToJobResponseDetail(data invitation.JobCore) JobDetailResponse {
	return JobDetailResponse{
		ID:           data.ID,
		Title:        data.Title,
		Description:  data.Description,
		RecruiterId:  data.RecruiterId,
		Requirements: ToRequirementsResponse(data.Requirements),
	}
}

func ToRequirementsResponse(data []invitation.RequirementCore) []RequirementResponse {
	converted := []RequirementResponse{}
	for _, req := range data {
		temp := RequirementResponse{
			ID:          req.ID,
			Description: req.Description,
		}
		converted = append(converted, temp)
	}

	return converted
}

type InvitationUserResponse struct {
	ID          uint
	RecruiterID uint
	Status      string
	Job         JobResponse
}

type JobResponse struct {
	ID          int
	Title       string
	Description string
}

type InvitationJobResponse struct {
	ID          uint
	RecruiterID uint
	Status      string
	User        UserResponse
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

func ToInvitationUserResponse(data invitation.InvitationCore) InvitationUserResponse {
	return InvitationUserResponse{
		ID:          data.ID,
		RecruiterID: data.RecruiterID,
		Status:      data.Status,
		Job:         ToJobResponse(data.Job),
	}
}

func ToJobResponse(data invitation.JobCore) JobResponse {
	return JobResponse{
		ID:          data.ID,
		Title:       data.Title,
		Description: data.Description,
	}
}

func ToInvitationUserResponseList(data []invitation.InvitationCore) []InvitationUserResponse {
	converted := []InvitationUserResponse{}
	for _, inv := range data {
		converted = append(converted, ToInvitationUserResponse(inv))
	}

	return converted
}

func ToInvitationJobResponse(data invitation.InvitationCore) InvitationJobResponse {
	return InvitationJobResponse{
		ID:          data.ID,
		RecruiterID: data.RecruiterID,
		Status:      data.Status,
		User:        ToUserResponse(data.User),
	}
}

func ToUserResponse(data invitation.UserCore) UserResponse {
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

func ToInvitationJobResponseList(data []invitation.InvitationCore) []InvitationJobResponse {
	converted := []InvitationJobResponse{}
	for _, inv := range data {
		converted = append(converted, ToInvitationJobResponse(inv))
	}

	return converted
}
