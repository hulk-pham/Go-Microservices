package user

import (
	"hulk/go-webservice/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api

// PingExample godoc
// @Summary List User
// @Schemes
// @Description list all user
// @Tags api
// @Accept json
// @Produce json
// @Success 200 {array} User
// @Router /user [get]
func GetListUserAction(c *gin.Context) {
	var users []User
	common.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func CreateUserAction(c *gin.Context) {
	var users []User
	common.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}
