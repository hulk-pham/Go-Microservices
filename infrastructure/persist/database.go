package persist

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"hulk/go-webservice/infrastructure/config"
	"io/ioutil"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func InitDB() *gorm.DB {
	config := config.AppConfig()

	isTLS := false

	if config.DBUseSSL == "true" {
		isTLS = true
		rootCertPool := x509.NewCertPool()
		pem, err := ioutil.ReadFile(config.DBCAFile)
		if err != nil {
			log.Fatal(err)
		}
		if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
			log.Fatal("Failed to append PEM.")
		}
		mysql.RegisterTLSConfig("custom", &tls.Config{RootCAs: rootCertPool})
	}

	params := make(map[string]string)
	params["parseTime"] = "true"

	cfg := mysql.Config{
		User:                 config.DBUsername,
		Passwd:               config.DBPassword,
		Addr:                 fmt.Sprintf("%s:%s", config.DBAddress, config.DBPort),
		Net:                  "tcp",
		DBName:               config.DBDatabase,
		Loc:                  time.Local,
		AllowNativePasswords: true,
		Params:               params,
	}

	if isTLS == true {
		cfg.TLSConfig = "custom"
	}

	str := cfg.FormatDSN()

	db, err := gorm.Open(gormMysql.Open(str), &gorm.Config{
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
