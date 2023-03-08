package user

import (
	"fmt"
	"hulk/go-webservice/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary List User
// @Schemes
// @Description list all user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} common.JSONResult{data=[]User}
// @Router /user [get]
func GetListUserAction(c *gin.Context) {
	fmt.Print(c.Get("CurrentUser"))
	var users []User
	common.DB.Find(&users)
	c.JSON(http.StatusOK, common.JSONResult{Code: 200, Message: "Ok", Data: users})
}

// PingExample godoc
// @Summary Create User
// @Param request body CreateUserRequest true "body params"
// @Description create user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} common.JSONResult{data=User}
// @Router /user [post]
func CreateUserAction(c *gin.Context) {
	var request CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	duplicated := User{Email: request.Email}
	if r := common.DB.Where(&User{Email: request.Email}).First(&duplicated); r.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already has been taken"})
		return
	}

	var user User
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
	fmt.Print(passwordHashed)
	user.Password = passwordHashed

	common.DB.Create(&user)

	c.JSON(http.StatusOK, common.JSONResult{Code: 200, Message: "Ok", Data: user})
}
