package routes

import (
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/hello", handlers.HelloHandler)

	router.GET("/users", handlers.ListUsers)
	router.POST("/users", handlers.CreateUser)

	router.GET("/books", handlers.ListBooks)

	router.POST("/sessions", handlers.SessionStart)

	return router
}
