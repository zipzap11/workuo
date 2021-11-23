package response

import (
	"net/http"
	"workuo/features/recruiter"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string      `json: "message"`
	Data    interface{} `json: "data"`
}

type RecruiterLoginResponse struct {
	ID      uint   `json: "id"`
	Company string `json: "company"`
	Address string `json: "address"`
	Token   string `json: "token"`
}

type RecruiterResponse struct {
	ID      uint   `json: "id"`
	Company string `json: "company"`
	Address string `json: "address"`
	Bio     string `json: "bio`
}

func NewSuccessResponse(e echo.Context, msg string, data interface{}) error {
	return e.JSON(http.StatusOK, Response{
		Message: msg,
		Data:    data,
	})
}

func ToRecruiterLoginResponse(data recruiter.RecruiterCore) RecruiterLoginResponse {
	return RecruiterLoginResponse{
		ID:      data.ID,
		Company: data.Company,
		Address: data.Address,
		Token:   data.Token,
	}
}

func ToRecruiterResponse(data recruiter.RecruiterCore) RecruiterResponse {
	return RecruiterResponse{
		ID:      data.ID,
		Company: data.Company,
		Address: data.Address,
		Bio:     data.Bio,
	}
}

func ToRecruiterResponseList(data []recruiter.RecruiterCore) []RecruiterResponse {
	convertedRec := []RecruiterResponse{}
	for _, rec := range data {
		convertedRec = append(convertedRec, ToRecruiterResponse(rec))
	}

	return convertedRec
}
