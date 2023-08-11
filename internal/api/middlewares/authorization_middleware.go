package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/gabrielmrts/mybooks-golang-api/internal/api/enums"
	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware(c *gin.Context) {
	var userData UserDataContext

	if user, exists := c.Get("user"); exists {
		if data, ok := user.(UserDataContext); ok {
			userData = data
		}
	}

	if strings.HasPrefix(c.Request.URL.Path, "/private/admin") {
		log.Printf("ROLE: %d", userData.Id)
		if userData.Role != enums.ROLES.ADMIN {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not allowed to access this endpoint"})
		}
	}

	c.Next()
}
