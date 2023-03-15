package middleware

import (
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	configCors := cors.DefaultConfig()
	configCors.AllowOrigins = []string{"*"}
	return cors.New(configCors)
}
