package handlers

import (
	"net/http"

	"github.com/gabrielmrts/mybooks-golang-api/internal/api/factories"
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/helpers"
	"github.com/gin-gonic/gin"
)

type SessionRequestBody struct {
	Email    string
	Password string
}

func SessionStart(c *gin.Context) {
	var requestBody SessionRequestBody
	usersRepository := factories.GetUsersRepository

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing email or password param"})
		return
	}

	passwordHash := helpers.GeneratePasswordHash(requestBody.Password)
	user, err := usersRepository().FindByEmailAndPassword(requestBody.Email, passwordHash)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "invalid credentials"})
		return
	}

	token, _ := helpers.GenerateJWT(user.ID)

	c.JSON(http.StatusCreated, gin.H{"bearer": token})
}
