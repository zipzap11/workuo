package application

import "time"

type ApplicationCore struct {
	ID        uint
	UserID    uint
	JobID     uint
	Status    string
	AppliedAt time.Time
}

type Repository interface {
}

type Service interface {
}
