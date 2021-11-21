package request

import "workuo/features/recruiter"

type RecruiterRequest struct {
	Company string `json: "company"`
	Address string `json: "address"`
	Bio     string `json: "bio"`
}

func ToCore(data RecruiterRequest) recruiter.RecruiterCore {
	return recruiter.RecruiterCore{
		Company: data.Company,
		Address: data.Address,
		Bio:     data.Bio,
	}
}
