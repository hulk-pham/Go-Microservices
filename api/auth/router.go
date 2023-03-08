package auth

import "github.com/gin-gonic/gin"

func Router(router *gin.RouterGroup) {
	userRouter := router.Group("/auth")

	userRouter.POST("/login", LoginAction)
}
