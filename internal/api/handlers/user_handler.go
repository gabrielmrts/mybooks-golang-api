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

// ListUsers lists all existing users
//
// @Summary      List users
// @Tags         Users
// @Produce      json
// @Success      200              {array}   models.Account
// @failure      400
// @failure      401
// @Security Bearer
// @Router		 /private/admin/users [get]
func ListUsers(c *gin.Context) {
	usersRepository := factories.GetUsersRepository()
	users, err := usersRepository.List()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to list users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// CreateUser
//
// @Summary      Create account
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      201
// @failure      400
// @failure      401
// @Param		 body	body	CreateUserRequestBody	true "request example"
// @Router		 /public/users [post]
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

	usersRepository := factories.GetUsersRepository()
	accountsRepository := factories.GetAccountsRepository()

	var user = models.User{
		Name: requestBody.Name,
	}
	var account = models.Account{
		Email:    requestBody.Email,
		Password: helpers.GeneratePasswordHash(requestBody.Password),
	}

	if _, err := usersRepository.FindByEmail(account.Email); err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "email already used"})
		return
	}

	if err := accountsRepository.Create(&account, &user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "error creating user"})
		return
	}

	c.AbortWithStatusJSON(http.StatusCreated, gin.H{"message": "user created with success"})
}
