package data

import (
	"workuo/features/recruiter"

	"gorm.io/gorm"
)

type Recruiter struct {
	gorm.Model
	Company  string
	Bio      string
	Address  string
	Email    string
	Password string
}

func FromCore(data recruiter.RecruiterCore) Recruiter {
	return Recruiter{
		Company:  data.Company,
		Bio:      data.Bio,
		Address:  data.Address,
		Email:    data.Email,
		Password: data.Password,
	}
}