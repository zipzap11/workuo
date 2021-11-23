package job

import (
	"time"
)

type JobCore struct {
	ID           int
	Title        string
	Description  string
	RecruiterId  int
	Company      string
	Requirements []RequirementCore
	Created_at   time.Time
	Updated_at   time.Time
}

type RequirementCore struct {
	ID          uint
	JobId       uint
	Description string
}

type Service interface {
	CreateJobPost(data JobCore) (err error)
	GetJobPost(data JobCore) ([]JobCore, error)
	GetJobPostById(data JobCore) (JobCore, error)
	// UpdateJobPost(data JobCore) (resp JobCore, err error)
	// DeleteJobPost(id int) (err error)
	// GetJobPostByTitle(title string) (resp []JobCore, err error)
	// GetJobPostByCompanyName(company string) (resp []JobCore, err error)
}

type Data interface {
	InsertData(data JobCore) (err error)
	GetJobData(data JobCore) ([]JobCore, error)
	GetJobDataById(data JobCore) (JobCore, error)
	// SelectAllData() (resp []JobCore, err error)
	// UpdatedData(data JobCore) (resp JobCore, err error)
	// DeleteData(id int) (err error)
	// SelectDataByTitle(title string) (resp []JobCore, err error)
	// SelectDataByCompany(company string) (resp []JobCore, err error)
}
