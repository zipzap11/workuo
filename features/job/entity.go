package job

import (
	"time"
	"workuo/features/job/presentation/request"
)

type JobCore struct {
	ID           int
	Title        string
	Description  string
	Recruiter_id int
	Requirements []RequirementCore
	Created_at   time.Time
	Updated_at   time.Time
}

type RequirementCore struct {
	ID          uint
	JobID       uint
	Description string
}

type Service interface {
	CreateJobPost(data request.Job) (err error)
	// UpdateJobPost(data JobCore) (resp JobCore, err error)
	// DeleteJobPost(id int) (err error)
	// GetAllJobPost() (resp []JobCore, err error)
	// GetJobPostByTitle(title string) (resp []JobCore, err error)
	// GetJobPostByCompanyName(company string) (resp []JobCore, err error)
	// GetJobById(id int) (resp JobCore, err error)
}

type Data interface {
	InsertData(data request.Job) (err error)
	// SelectAllData() (resp []JobCore, err error)
	// UpdatedData(data JobCore) (resp JobCore, err error)
	// DeleteData(id int) (err error)
	// SelectDataByTitle(title string) (resp []JobCore, err error)
	// SelectDataByCompany(company string) (resp []JobCore, err error)
	// SelectDataById(id int) (resp JobCore, err error)
}
