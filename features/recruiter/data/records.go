package data

import "gorm.io/gorm"

type Recruiter struct {
	gorm.Model
	Company string
	Bio     string
	Address string
}
