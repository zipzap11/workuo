package driver

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/training?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigrate(DB)
	return DB
}

func initMigrate(DB *gorm.DB) {
	DB.AutoMigrate()
}
