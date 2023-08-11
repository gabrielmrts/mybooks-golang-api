package middlewares

import (
	"net/http"
	"strings"

	"github.com/gabrielmrts/mybooks-golang-api/internal/api/enums"
	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware(c *gin.Context) {

	userRole, _ := c.Get("user_role")

	if strings.HasPrefix(c.Request.URL.Path, "/private/admin") {
		if userRole != enums.ROLES.ADMIN {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not allowed to access this endpoint"})
		}
	}

	c.Next()
}
