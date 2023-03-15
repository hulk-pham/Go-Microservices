package main

import (
	"hulk/go-webservice/common"
	"hulk/go-webservice/domain/entities"

	"hulk/go-webservice/infrastructure/config"
	"hulk/go-webservice/infrastructure/persist"

	"hulk/go-webservice/presentation/graph"
	"hulk/go-webservice/presentation/http"
	"hulk/go-webservice/presentation/http/docs"
	"hulk/go-webservice/presentation/http/middleware"
	"hulk/go-webservice/presentation/realtime"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
		c.Header("Content-Type", "application/json")
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	config := config.AppConfig()
	docs.SwaggerInfo.BasePath = "/api"

	if config.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	persist.InitDB()
	common.InitCacheService()
	realtime.InitRoomManager()

	persist.DB.AutoMigrate(&entities.User{})

	r := http.InitRouter()
	r.Use(middleware.CORSMiddleware())

	r.GET("/ws", realtime.WShandler())
	r.POST("/query", graphqlHandler())
	r.GET("/graphql-playground", playgroundHandler())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":" + config.AppPort)

}
