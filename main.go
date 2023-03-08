package main

import (
	"hulk/go-webservice/api/user"
	"hulk/go-webservice/common"
	"hulk/go-webservice/configs"

	docs "hulk/go-webservice/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	config, err := configs.LoadAppConfig(".")
	if err != nil {
		panic("failed to local config")
	}

	var db = common.Init()
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&user.User{})

	r := gin.Default()

	apiGroup := r.Group("/api")
	user.Router(apiGroup)

	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":" + config.AppPort)
}
