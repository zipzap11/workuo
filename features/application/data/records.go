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
	AppliedAt time.Time
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
	}
}

func ToCoreList(data []Application) []application.ApplicationCore {
	convertedData := []application.ApplicationCore{}
	for _, app := range data {
		convertedData = append(convertedData, ToCore(app))
	}
	return convertedData
}
