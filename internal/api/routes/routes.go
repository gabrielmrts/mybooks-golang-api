package routes

import (
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/hello", handlers.HelloHandler)

	return router
}
