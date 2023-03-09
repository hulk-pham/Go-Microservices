package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func InitDB() *gorm.DB {
	config := AppConfig()

	db, err := gorm.Open(mysql.Open(config.DBSource), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
