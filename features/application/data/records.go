package data

import (
	"time"
	"workuo/features/application"

	"gorm.io/gorm"
)

type Application struct {
	gorm.Model
	UserID    uint
	JobID     uint
	Status    string
	Job       Job
	User      User
	AppliedAt time.Time
}

type Job struct {
	gorm.Model
	Title        string
	Description  string
	RecruiterId  int
	Requirements []Requirement
}

type Requirement struct {
	gorm.Model
	JobID       uint
	Description string
}

type User struct {
	gorm.Model
	Name    string
	Dob     time.Time
	Gender  string
	Address string
	Title   string
	Bio     string
}

func (j *Job) toCore() application.JobCore {
	convertedRequirement := []application.RequirementCore{}
	for _, req := range j.Requirements {
		convertedRequirement = append(convertedRequirement, req.toCore())
	}
	return application.JobCore{
		ID:           int(j.ID),
		Title:        j.Title,
		Description:  j.Description,
		RecruiterId:  j.RecruiterId,
		Requirements: convertedRequirement,
		Created_at:   j.CreatedAt,
		Updated_at:   j.UpdatedAt,
	}
}

func (u User) toCore() application.UserCore {
	return application.UserCore{
		ID:      u.ID,
		Name:    u.Name,
		Dob:     u.Dob,
		Gender:  u.Gender,
		Title:   u.Title,
		Address: u.Address,
		Bio:     u.Bio,
	}
}

func (r *Requirement) toCore() application.RequirementCore {
	return application.RequirementCore{
		ID:          r.ID,
		JobId:       r.JobID,
		Description: r.Description,
	}
}

func toCoreList(jobs []Job) []application.JobCore {
	var convertedData []application.JobCore
	for _, job := range jobs {
		convertedData = append(convertedData, job.toCore())
	}

	return convertedData
}

func ToApplicationRecord(data application.ApplicationCore) Application {
	return Application{
		UserID:    data.UserID,
		JobID:     data.JobID,
		Status:    data.Status,
		AppliedAt: data.AppliedAt,
	}
}

func ToCore(data Application) application.ApplicationCore {
	return application.ApplicationCore{
		ID:        data.ID,
		UserID:    data.UserID,
		JobID:     data.JobID,
		Status:    data.Status,
		AppliedAt: data.AppliedAt,
		Job:       data.Job.toCore(),
		User:      data.User.toCore(),
	}
}

func ToCoreList(data []Application) []application.ApplicationCore {
	convertedData := []application.ApplicationCore{}
	for _, app := range data {
		convertedData = append(convertedData, ToCore(app))
	}
	return convertedData
}
