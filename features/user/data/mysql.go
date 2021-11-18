package data

import (
	"workuo/features/user"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	DB *gorm.DB
}

func NewMydsqlRepository(DB *gorm.DB) user.Repository {
	return &mysqlUserRepository{}
}

func (mr *mysqlUserRepository) InsertData(data user.UserCore) error {
	err := mr.DB.Create(toUserRecord(data))
	if err != nil {
		return err.Error
	}

	return nil
}
