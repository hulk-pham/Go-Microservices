package auth

import (
	"hulk/go-webservice/application/modules/user/commands"
	"hulk/go-webservice/domain/entities"
	"hulk/go-webservice/presentation/http/base"
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
// @Success 200 {object} base.JSONResult{data=string}
// @Router /auth/login [post]
func LoginAction(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := commands.LoginCommand(commands.LoginRequestDto(request))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, base.JSONResult{Code: 200, Message: "Ok", Data: token})
}

// PingExample godoc
// @Summary Sign Up
// @Param request body SignUpRequest true "body params"
// @Description create user
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} base.JSONResult{data=entities.User}
// @Router /auth/signup [post]
// 2006-01-02T15:04:05.000Z
func SignupAction(c *gin.Context) {
	var request SignUpRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user entities.User
	user, err := commands.CreateUserCommand(commands.CreateUserDto(request))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, base.JSONResult{Code: 200, Message: "Ok", Data: user})
}
