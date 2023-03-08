package common

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.Replace(c.Request.Header["Authorization"][0], "Bearer ", "", -1)

		userClaim, err := ValidateJWT(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Invalid token"})
			return
		}

		c.Set("CurrentUser", userClaim)

		c.Next()
	}
}
