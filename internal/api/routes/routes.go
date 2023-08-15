package routes

import (
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/handlers"
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/middlewares"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	// TO allow CORS
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(CORS())

	privateGroup := router.Group("/private")
	privateGroup.Use(middlewares.AuthenticationMiddleware)
	privateGroup.Use(middlewares.AuthorizationMiddleware)
	privateGroup.POST("/books", handlers.CreateBook)
	privateGroup.GET("/sessions/me", handlers.SessionsMe)

	publicGroup := router.Group("/public")
	publicGroup.GET("/hello", handlers.HelloHandler)
	publicGroup.GET("/books", handlers.ListBooks)
	publicGroup.POST("/users", handlers.CreateUser)
	publicGroup.POST("/users/email", handlers.VerifyEmail)
	publicGroup.POST("/sessions", handlers.SessionStart)
	publicGroup.POST("/users/password/recover", handlers.PasswordRecoverLinkEmailHandler)
	publicGroup.PATCH("/users/password", handlers.ResetPassword)

	adminGroup := privateGroup.Group("/admin")
	adminGroup.GET("/users", handlers.ListUsers)

	return router
}
