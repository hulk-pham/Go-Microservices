package main

import (
	"hulk/go-webservice/api"
	"hulk/go-webservice/core/model"
	"hulk/go-webservice/common"
	"hulk/go-webservice/graph"

	docs "hulk/go-webservice/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

// Defining the Graphql handler
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
	if config.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	var db = common.InitDB()
	db.AutoMigrate(&model.User{})

	r := api.InitRouter()

	r.POST("/query", graphqlHandler())
	r.GET("/graphql-playground", playgroundHandler())

	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":" + config.AppPort)
}
