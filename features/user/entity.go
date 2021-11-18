package user

import "time"

type UserCore struct {
	Id      uint
	Name    string
	Dob     time.Time
	Gender  string
	Address string
	Title   string
	Bio     string
}
