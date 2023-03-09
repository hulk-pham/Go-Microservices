package main

import (
	"hulk/go-webservice/api"
	"hulk/go-webservice/api/user"
	"hulk/go-webservice/common"

	docs "hulk/go-webservice/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config := common.AppConfig()
	if config.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	var db = common.InitDB()
	db.AutoMigrate(&user.User{})

	r := api.InitRouter()

	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":" + config.AppPort)
}
