package data

import (
	"fmt"
	"workuo/features/user"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	DB *gorm.DB
}

func NewMysqlUserRepository(DB *gorm.DB) user.Repository {
	return &mysqlUserRepository{DB}
}

func (mr *mysqlUserRepository) InsertData(data user.UserCore) error {
	fmt.Println("data in repository ======", data)
	recordData := toUserRecord(data)
	fmt.Println("data in repository converted ======", recordData)
	err := mr.DB.Create(&recordData)
	if err != nil {
		return err.Error
	}

	return nil
}
