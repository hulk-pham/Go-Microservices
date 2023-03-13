package main

import (
	"hulk/go-webservice/api"
	"hulk/go-webservice/common"
	"hulk/go-webservice/core/model"
	"hulk/go-webservice/graph"
	"hulk/go-webservice/realtime"

	docs "hulk/go-webservice/docs"

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
	config := common.AppConfig()
	docs.SwaggerInfo.BasePath = "/api"

	if config.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	common.InitDB()
	common.InitCacheService()
	realtime.InitRoomManager()

	common.DB.AutoMigrate(&model.User{})

	r := api.InitRouter()
	r.Use(common.CORSMiddleware())

	r.GET("/ws", realtime.WShandler())
	r.POST("/query", graphqlHandler())
	r.GET("/graphql-playground", playgroundHandler())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":" + config.AppPort)

}
