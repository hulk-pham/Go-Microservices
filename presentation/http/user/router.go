package user

import (
	"hulk/go-webservice/presentation/http/middleware"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	userRouter := router.Group("/user")
	userRouter.Use(middleware.AuthMiddleware())

	userRouter.GET("/", GetListUserAction)
	userRouter.POST("/", CreateUserAction)
	userRouter.POST(":id/avatar-upload", UserUpdateAvatarAction)
}
