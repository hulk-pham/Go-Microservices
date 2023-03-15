package persist

import (
	"hulk/go-webservice/infrastructure/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func InitDB() *gorm.DB {
	config := config.AppConfig()

	db, err := gorm.Open(mysql.Open(config.DBSource), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
