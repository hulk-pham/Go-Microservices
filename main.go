package main

import (
	"hulk/go-webservice/api/auth"
	"hulk/go-webservice/api/user"
	"hulk/go-webservice/common"
	"hulk/go-webservice/configs"

	docs "hulk/go-webservice/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func main() {
	config := configs.AppConfig()

	var db = common.Init()
	db.AutoMigrate(&user.User{})

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.Static("/public", "./public")

	// @BasePath /api
	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization
	apiGroup := r.Group("/api")

	auth.Router(apiGroup)
	user.Router(apiGroup)

	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":" + config.AppPort)
}
