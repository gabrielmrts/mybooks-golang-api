package handlers

import (
	"log"
	"net/http"

	"github.com/gabrielmrts/mybooks-golang-api/internal/api/factories"
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/models"
	"github.com/gin-gonic/gin"
)

type CreateUserRequestBody struct {
	Name     string
	Email    string
	Password string
}

func ListUsers(c *gin.Context) {
	usersRepository := factories.GetUsersRepository()
	users, err := usersRepository.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var requestBody CreateUserRequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing parameter"})
		return
	}

	var user = models.User{Name: requestBody.Name, Email: requestBody.Email, Password: requestBody.Email}

	usersRepository := factories.GetUsersRepository()
	if err := usersRepository.Create(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error creating user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created with success"})
}
