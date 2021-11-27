package application

import (
	"time"
)

type ApplicationCore struct {
	ID        uint
	UserID    uint
	JobID     uint
	Status    string
	AppliedAt time.Time
}

type Repository interface {
	ApplyJob(ApplicationCore) error
	GetApplicationByUserID(int) ([]ApplicationCore, error)
	RejectApplication(int) error
	AcceptApplication(int) error
}

type Service interface {
	ApplyJob(ApplicationCore) error
	GetApplicationByUserID(int) ([]ApplicationCore, error)
	RejectApplication(int) error
	AcceptApplication(int) error
}
