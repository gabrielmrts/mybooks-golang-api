package handlers

import (
	"net/http"

	"github.com/gabrielmrts/mybooks-golang-api/internal/api/factories"
	"github.com/gin-gonic/gin"
)

func ListUsers(c *gin.Context) {
	usersRepository := factories.GetUsersRepository()
	users, err := usersRepository.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list users"})
		return
	}

	c.JSON(http.StatusOK, users)
}
