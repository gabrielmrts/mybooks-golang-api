package middlewares

import (
	"net/http"

	"github.com/gabrielmrts/mybooks-golang-api/internal/api/helpers"
	"github.com/gin-gonic/gin"
)

type UserDataContext struct {
	Id   uint
	Exp  int64
	Role string
}

func AuthenticationMiddleware(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.Next()
	}

	token := c.GetHeader("authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token not provided"})
		return
	}

	var userDataContext = UserDataContext{}

	claims, err := helpers.VerifyToken(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	userDataContext.Role = claims.Role
	userDataContext.Id = claims.Id
	userDataContext.Exp = claims.Exp

	c.Set("user", userDataContext)
	c.Next()
}
