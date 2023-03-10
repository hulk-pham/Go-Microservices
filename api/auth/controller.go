package auth

import (
	"hulk/go-webservice/core/model"
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
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound model.User
	if r := common.DB.Where(&model.User{Email: request.Email}).First(&userFound); r.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not exist"})
		return
	}

	if err := common.CheckPassword(request.Password, userFound.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password does not match"})
		return
	}

	token, err := common.GenerateJWT(common.UserClaim{Username: userFound.FirstName + "" + userFound.LastName, Id: userFound.ID})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable generate token"})
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
// @Success 200 {object} common.JSONResult{data=model.User}
// @Router /auth/signup [post]
func SignupAction(c *gin.Context) {
	var request SignUpRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	duplicated := model.User{Email: request.Email}
	if r := common.DB.Where(&model.User{Email: request.Email}).First(&duplicated); r.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already has been taken"})
		return
	}

	var user model.User
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Email = request.Email
	user.Address = request.Address
	user.Hobby = request.Hobby
	user.PhoneNumber = request.PhoneNumber
	user.Dob = request.Dob
	passwordHashed, err := common.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to hass password"})
		return
	}
	user.Password = passwordHashed

	common.DB.Create(&user)

	c.JSON(http.StatusOK, common.JSONResult{Code: 200, Message: "Ok", Data: user})
}
