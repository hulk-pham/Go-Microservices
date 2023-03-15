package auth

import (
	"hulk/go-webservice/application/modules/user/commands"
	"hulk/go-webservice/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary Login User
// @Param request body LoginRequest true "body params"
// @Description login
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} common.JSONResult{data=string}
// @Router /auth/login [post]
func LoginAction(c *gin.Context) {
	var request LoginRequest

	token, err := commands.LoginCommand(commands.LoginRequestDto(request))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.JSONResult{Code: 200, Message: "Ok", Data: token})
}

// PingExample godoc
// @Summary Sign Up
// @Param request body SignUpRequest true "body params"
// @Description create user
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} common.JSONResult{data=entities.User}
// @Router /auth/signup [post]
func SignupAction(c *gin.Context) {
	var request SignUpRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := commands.CreateUserCommand(commands.CreateUserDto(request))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.JSONResult{Code: 200, Message: "Ok", Data: user})
}
