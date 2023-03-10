package user

import (
	"fmt"
	"hulk/go-webservice/common"
	"hulk/go-webservice/core/model"
	"hulk/go-webservice/core/dto"
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
// @Security ApiKeyAuth
// @Router /user [get]
func GetListUserAction(c *gin.Context) {
	fmt.Print(c.Get("CurrentUser"))
	var users []model.User
	common.DB.Find(&users)
	c.JSON(http.StatusOK, common.JSONResult{Code: 200, Message: "Ok", Data: users})
}

// PingExample godoc
// @Summary Create User
// @Param request body dto.CreateUserRequest true "body params"
// @Description create user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} common.JSONResult{data=model.User}
// @Security ApiKeyAuth
// @Router /user [post]
func CreateUserAction(c *gin.Context) {
	var request dto.CreateUserRequest
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
	fmt.Print(passwordHashed)
	user.Password = passwordHashed

	common.DB.Create(&user)

	c.JSON(http.StatusOK, common.JSONResult{Code: 200, Message: "Ok", Data: user})
}

// PingExample godoc
// @Summary User Update Avatar
// @Param id path int true "User ID"
// @Param file formData file true "file"
// @Description create user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} common.JSONResult{data=model.User}
// @Security ApiKeyAuth
// @Router /user/{id}/avatar-upload [post]
func UserUpdateAvatarAction(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Required file"})
		return
	}

	userID := c.Params.ByName("id")
	var user model.User
	if r := common.DB.First(&user, userID); r.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	common.UploadLocal(c, file)
	fileUploadedPath, err := common.UploadS3(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Upload file fail"})
		return
	}
	common.RemoveFile(file.Filename)

	user.Avatar = fileUploadedPath
	common.DB.Save(&user)

	c.JSON(http.StatusOK, common.JSONResult{Code: 200, Message: "Ok", Data: user})
}
