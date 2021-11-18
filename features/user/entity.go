package user

import "time"

type UserCore struct {
	Id      uint
	Name    string
	dob     time.Time
	Gender  string
	Address string
	Title   string
	Bio     string
}
