package data

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string    `json: "name"`
	Address string    `json: "address"`
	Bio     string    `json: "name"`
	dob     time.Time `json: "dob"`
	Title   string    `json: "title"`
	Gender  string    `json: "gender"`
}
