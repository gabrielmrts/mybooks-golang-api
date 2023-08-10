package middlewares

import (
	"net/http"
	"strings"

	"github.com/gabrielmrts/mybooks-golang-api/internal/api/enums"
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/helpers"
	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware(c *gin.Context) {
	userData := helpers.JWTClaims{}

	if user, exists := c.Get("user"); exists {
		if data, ok := user.(helpers.JWTClaims); ok {
			userData = data
		}
	}

	if strings.HasPrefix(c.Request.URL.Path, "/private/admin") {
		if userData.Role != enums.ROLES.ADMIN {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not allowed to access this endpoint"})
		}
	}

	c.Next()
}
