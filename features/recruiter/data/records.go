package data

import (
	"workuo/features/recruiter"

	"gorm.io/gorm"
)

type Recruiter struct {
	gorm.Model
	Company string
	Bio     string
	Address string
}

func fromCore(data recruiter.RecruiterCore) Recruiter {
	return Recruiter{
		Company: data.Company,
		Bio:     data.Bio,
		Address: data.Address,
	}
}
