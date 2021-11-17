package request

type Job struct {
	Title        string   `json: "title"`
	Description  string   `json: "description"`
	RecruiterID  int      `json: "recruiter_id"`
	Requirements []string `json: requirements`
}
