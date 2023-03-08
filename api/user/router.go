package user

import "github.com/gin-gonic/gin"

func Router(router *gin.RouterGroup) {
	userRouter := router.Group("/user")

	userRouter.GET("/", GetListUserAction)
	userRouter.POST("/", CreateUserAction)
}
