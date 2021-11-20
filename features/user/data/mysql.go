package data

import (
	"errors"
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

func (mr *mysqlUserRepository) GetData() ([]user.UserCore, error) {
	var users []User

	err := mr.DB.Preload("Skillsets").Preload("Experiences").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return toUserCoreList(users), nil
}

func (mr *mysqlUserRepository) CheckUser(data user.UserCore) (user.UserCore, error) {
	var userData User
	err := mr.DB.Where("email = ? and password = ?", data.Email, data.Password).First(&userData).Error

	if userData.Name == "" && userData.ID == 0 {
		return user.UserCore{}, errors.New("no existing user")
	}
	if err != nil {
		return user.UserCore{}, err
	}

	return toUserCore(userData), nil
}

func (mr *mysqlUserRepository) GetDataById(data user.UserCore) (user.UserCore, error) {
	var userData User
	err := mr.DB.First(&userData, data.Id).Error

	if userData.Name == "" && userData.ID == 0 {
		return user.UserCore{}, errors.New("no existing user")
	}
	if err != nil {
		return user.UserCore{}, err
	}

	return toUserCore(userData), nil
}
