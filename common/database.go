package common

import (
	"hulk/go-webservice/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
	config, err := configs.LoadAppConfig(".")

	if err != nil {
		panic("failed to local config")
	}

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
