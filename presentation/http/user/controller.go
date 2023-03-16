package user

import (
	"encoding/json"
	"hulk/go-webservice/application/modules/user/commands"
	"hulk/go-webservice/application/modules/user/queries"
	"hulk/go-webservice/common"
	"hulk/go-webservice/domain/entities"
	"hulk/go-webservice/presentation/http/base"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary List User
// @Schemes
// @Description list all user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} base.JSONResult{data=[]entities.User}
// @Security ApiKeyAuth
// @Router /user [get]
func GetListUserAction(c *gin.Context) {
	var users []entities.User
	cacheValue, _ := common.CacheInstance.Get("golang:users")
	if cacheValue != "" {
		json.Unmarshal([]byte(cacheValue), &users)
		c.JSON(http.StatusOK, base.JSONResult{Code: 200, Message: "Ok", Data: users})
		return
	}
	users = queries.GetAllUserQuery()
	usersStr, _ := json.Marshal(&users)
	common.CacheInstance.Set("golang:users", string(usersStr), time.Minute*5)
	c.JSON(http.StatusOK, base.JSONResult{Code: 200, Message: "Ok", Data: users})
}

// PingExample godoc
// @Summary Create User
// @Param request body commands.CreateUserDto true "body params"
// @Description create user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} base.JSONResult{data=entities.User}
// @Security ApiKeyAuth
// @Router /user [post]
func CreateUserAction(c *gin.Context) {
	var request commands.CreateUserDto
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := commands.CreateUserCommand(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, base.JSONResult{Code: 200, Message: "Ok", Data: user})
}

// PingExample godoc
// @Summary User Update Avatar
// @Param id path int true "User ID"
// @Param file formData file true "file"
// @Description create user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} base.JSONResult{data=entities.User}
// @Security ApiKeyAuth
// @Router /user/{id}/avatar-upload [post]
func UserUpdateAvatarAction(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Required file"})
		return
	}
	common.UploadLocal(c, file)
	user, err := commands.UpdateUserAvatarCommand(c.Params.ByName("id"), file)
	if err != nil {
		common.RemoveFile(file.Filename)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, base.JSONResult{Code: 200, Message: "Ok", Data: user})
}
