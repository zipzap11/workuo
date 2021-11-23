package response

import "workuo/features/recruiter"

type RecruiterLoginResponse struct {
	ID      uint   `json: "id"`
	Company string `json: "company"`
	Address string `json: "address"`
	Token   string `json: "token"`
}

func ToRecruiterLoginResponse(data recruiter.RecruiterCore) RecruiterLoginResponse {
	return RecruiterLoginResponse{
		ID:      data.ID,
		Company: data.Company,
		Address: data.Address,
		Token:   data.Token,
	}
}
