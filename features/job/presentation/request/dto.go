package request

// import "workuo/features/job/data"

type Job struct {
	Title        string   `json: "title"`
	Description  string   `json: "description"`
	RecruiterID  int      `json: "recruiter_id"`
	Requirements []string `json: "requirements"`
}

// func (j *Job) toRecord() data.Job {
// 	convertedRequirements := []data.Requirement{}
// 	for _, req := range j.Requirements {
// 		convertedRequirements = append(convertedRequirements, data.Requirement{
// 			Description: req,
// 		})
// 	}
// 	return data.Job{
// 		Title: j.Title,
// 		Description: j.Description,
// 		Recruiter_id: j.RecruiterID,
// 		Requirements: convertedRequirements,
// 	}
// }
