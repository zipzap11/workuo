package request

import "workuo/features/recruiter"

type RecruiterRequest struct {
	Company  string `json: "company"`
	Address  string `json: "address"`
	Bio      string `json: "bio"`
	Email    string `json: "email"`
	Password string `json: "password"`
}

type RecruiterLogin struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

func ToCore(data RecruiterRequest) recruiter.RecruiterCore {
	return recruiter.RecruiterCore{
		Company:  data.Company,
		Address:  data.Address,
		Bio:      data.Bio,
		Email:    data.Email,
		Password: data.Password,
	}
}
