package routes

import (
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/handlers"
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	privateGroup := router.Group("/private")
	privateGroup.Use(middlewares.AuthenticationMiddleware)
	privateGroup.Use(middlewares.AuthorizationMiddleware)
	privateGroup.POST("/books", handlers.CreateBook)

	publicGroup := router.Group("/public")
	publicGroup.GET("/hello", handlers.HelloHandler)
	publicGroup.GET("/books", handlers.ListBooks)
	publicGroup.POST("/users", handlers.CreateUser)
	publicGroup.POST("/sessions", handlers.SessionStart)

	adminGroup := privateGroup.Group("/admin")
	adminGroup.GET("/users", handlers.ListUsers)

	return router
}
