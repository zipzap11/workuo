package data

import (
	"errors"
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
	recordData := toUserRecord(data)
	err := mr.DB.FirstOrCreate(&recordData)
	if err != nil {
		return err.Error
	}

	return nil
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
	err := mr.DB.Preload("Skillsets").Preload("Experiences").First(&userData, data.Id).Error

	if userData.Name == "" && userData.ID == 0 {
		return user.UserCore{}, errors.New("no existing user")
	}
	if err != nil {
		return user.UserCore{}, err
	}

	return toUserCore(userData), nil
}

func (mr *mysqlUserRepository) GetData(data user.UserCore) ([]user.UserCore, error) {
	var users []User

	titleFilter := "%" + data.Title + "%"

	var skillsetFilter string
	for i, skill := range data.Skillsets {
		if i == 0 {
			skillsetFilter = " AND skillsets.name in ("
		}
		if i == len(data.Skillsets)-1 {
			skillsetFilter += "'" + skill.Name + "')"
			break
		}
		skillsetFilter += "'" + skill.Name + "',"
	}

	err := mr.DB.Debug().Distinct("users.id", "users.name", "users.dob", "users.gender", "users.address", "users.title", "users.bio").
		Where("title LIKE ?", titleFilter).Preload("Skillsets").Preload("Experiences").
		Joins("inner JOIN user_skillsets ON user_skillsets.user_id = users.id inner JOIN skillsets ON skillsets.id = user_skillsets.skillset_id" + skillsetFilter).
		Find(&users).Error

	if err != nil {
		return nil, err
	}

	return toUserCoreList(users), nil
}
