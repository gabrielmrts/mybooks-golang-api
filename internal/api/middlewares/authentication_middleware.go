package middlewares

import (
	"net/http"

	"github.com/gabrielmrts/mybooks-golang-api/internal/api/helpers"
	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.Next()
	}

	token := c.GetHeader("authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token not provided"})
		return
	}

	claims, err := helpers.VerifyToken(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.Set("user", claims)
	c.Next()
}
