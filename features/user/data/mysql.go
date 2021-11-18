package data

import (
	"workuo/features/user"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	DB *gorm.DB
}

func NewMysqlUserRepository(DB *gorm.DB) user.Repository {
	return &mysqlUserRepository{}
}

func (mr *mysqlUserRepository) InsertData(data user.UserCore) error {
	recordData := toUserRecord(data)
	err := mr.DB.Create(&recordData)
	if err != nil {
		return err.Error
	}

	return nil
}
