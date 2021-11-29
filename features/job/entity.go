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
	GetJobPostById(id int) (JobCore, error)
	DeleteJobPost(data JobCore) (err error)
	UpdateJobPost(data JobCore) error
}

type Data interface {
	InsertData(data JobCore) (err error)
	GetJobData(data JobCore) ([]JobCore, error)
	GetJobDataById(id int) (JobCore, error)
	DeleteJobData(data JobCore) error
	UpdateJobData(data JobCore) error
}
