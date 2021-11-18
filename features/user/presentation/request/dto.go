package request

import "time"

type UserRequest struct {
	Name        string    `json: "name"`
	Address     string    `json: "address"`
	Dob         time.Time `json: "dob"`
	Gender      string    `json: "gender"`
	Bio         string    `json: "bio"`
	Title       string    `json: "title"`
	Skillsets   []string  `json: "skillsets"`
	Experiences []string  `json: "experiences`
}
