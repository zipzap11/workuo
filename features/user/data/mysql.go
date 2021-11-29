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

func (mr *mysqlUserRepository) InsertUserData(data user.UserCore) (int, error) {
	// recordData := toUserRecord(data)
	userData, _, _ := SeparateUserData(data)
	err := mr.DB.Create(&userData).Error

	if err != nil {
		return 0, err
	}

	return int(userData.ID), nil
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

func (mr *mysqlUserRepository) GetDataById(id int) (user.UserCore, error) {
	var userData User
	err := mr.DB.Preload("Skillsets").Preload("Experiences").First(&userData, id).Error

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

func (mr *mysqlUserRepository) UpdateUser(data user.UserCore) error {
	err := mr.DB.Debug().Model(&User{}).Where("id = ?", data.Id).Updates(User{
		Name:    data.Name,
		Dob:     data.Dob,
		Gender:  data.Gender,
		Address: data.Address,
		Title:   data.Title,
		Bio:     data.Bio,
		Email:   data.Email,
	}).Error
	if err != nil {
		return nil
	}

	return nil
}

func (mr *mysqlUserRepository) UpdateExperience(data user.ExperienceCore) error {
	err := mr.DB.Debug().Model(&Experience{}).Where("id = ?", data.Id).Updates(Experience{
		Title:       data.Title,
		Description: data.Description,
		StartDate:   data.StartDate,
		EndDate:     data.EndDate,
	}).Error
	if err != nil {
		return err
	}

	return nil
}

func (mr *mysqlUserRepository) CreateExperience(data user.ExperienceCore) error {
	record := ToExperienceRecord(data)
	err := mr.DB.Debug().Select("UserID", "Title", "Description", "StartDate", "EndDate").Create(&record).Error
	if err != nil {
		return err
	}
	return nil
}

func (mr *mysqlUserRepository) DeleteExperience(id int) error {
	err := mr.DB.Debug().Delete(&Experience{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (mr *mysqlUserRepository) CreateSkillset(data user.SkillsetCore) (int, error) {
	record := ToSkillsetRecord(data)
	record.ID = 0
	err := mr.DB.Debug().Where(Skillset{Name: data.Name}).FirstOrCreate(&record).Error

	if err != nil {
		return 0, err
	}
	return int(record.ID), nil
}

func (mr *mysqlUserRepository) AddUserSkillset(userId int, skillsetId int) error {
	err := mr.DB.Model(&UserSkillset{}).Create(UserSkillset{uint(userId), uint(skillsetId)}).Error
	if err != nil {
		return nil
	}

	return nil
}

func (mr *mysqlUserRepository) UpdateUserSkillset(userId int, skillsetId int, newSkillsetId int) error {
	err := mr.DB.Model(&UserSkillset{}).Where("user_id = ? AND skillset_id = ?", userId, skillsetId).Update("skillset_id", newSkillsetId).Error
	if err != nil {
		return nil
	}

	return nil
}

func (mr *mysqlUserRepository) DeleteUserSkillset(userId int, skillsetId int) error {
	err := mr.DB.Debug().Where("user_id = ? AND skillset_id = ?", userId, skillsetId).Delete(&UserSkillset{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (mr *mysqlUserRepository) GetUserSkillsets(data user.UserCore) ([]user.UserSkillsetCore, error) {
	var userSkillsets []UserSkillset
	err := mr.DB.Model(&UserSkillset{}).Where("user_id = ?", data.Id).Find(&userSkillsets).Error
	if err != nil {
		return nil, err
	}
	return ToUserSkillsetCore(userSkillsets), nil
}

func (mr *mysqlUserRepository) GetUserByEmail(email string) (bool, error) {
	var userModel User
	err := mr.DB.Where("email = ?", email).Find(&userModel).Error
	if err != nil {
		return false, err
	}
	if userModel.ID != 0 {
		return true, nil
	}
	return false, nil
}
