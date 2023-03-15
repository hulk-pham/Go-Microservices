package http

import (
	"hulk/go-webservice/presentation/http/auth"
	"hulk/go-webservice/presentation/http/user"

	"github.com/gin-gonic/gin"
)

// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func InitRouter() (r *gin.Engine) {
	r = gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.Static("/public", "./public")

	apiGroup := r.Group("/api")

	auth.Router(apiGroup)
	user.Router(apiGroup)

	return r
}
