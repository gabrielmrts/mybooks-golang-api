package handlers

import (
	"errors"
	"net/http"

	"github.com/gabrielmrts/mybooks-golang-api/internal/api/factories"
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/helpers"
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateUserRequestBody struct {
	Name     string `binding:"required,alphanum,min=4,max=255"`
	Email    string `binding:"required,email,min=3"`
	Password string `binding:"required,min=6,max=255"`
}

func ListUsers(c *gin.Context) {
	usersRepository := factories.GetUsersRepository()
	users, err := usersRepository.List()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to list users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var requestBody CreateUserRequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": out})
		}
		return
	}

	passwordHash := helpers.GeneratePasswordHash(requestBody.Password)
	var user = models.User{Name: requestBody.Name, Email: requestBody.Email, Password: passwordHash}

	usersRepository := factories.GetUsersRepository()

	if _, err := usersRepository.FindByEmail(user.Email); err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "email already used"})
		return
	}
	if err := usersRepository.Create(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "error creating user"})
		return
	}

	c.AbortWithStatusJSON(http.StatusCreated, gin.H{"message": "user created with success"})
}
