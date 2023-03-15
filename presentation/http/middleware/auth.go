package middleware

import (
	"hulk/go-webservice/common"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Request.Header["Authorization"]) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Require token"})
			return
		}

		token := strings.Replace(c.Request.Header["Authorization"][0], "Bearer ", "", -1)
		userClaim, err := common.ValidateJWT(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Invalid token"})
			return
		}

		c.Set("CurrentUser", userClaim)
		c.Next()
	}
}
