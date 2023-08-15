package handlers

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gabrielmrts/mybooks-golang-api/internal/api/factories"
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/helpers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// PasswordRecoverLinkEmail
//
// @Summary      Send password recover link to user email
// @Tags         Users
// @Accept       json
// @Success      201
// @failure      400
// @failure      401
// @Param		 body	body	handlers.PasswordRecoverLinkEmailHandler.request	true "body example"
// @Router		 /public/users/password/recover [post]
func PasswordRecoverLinkEmailHandler(c *gin.Context) {
	type request struct {
		Email string `json:"email" binding:"required,email,min=3"`
	}

	var requestBody = request{}
	mailService := factories.GetMailService()
	usersRepository := factories.GetUsersRepository()

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

	user, err := usersRepository.FindByEmail(requestBody.Email)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "email not found"})
	}

	expTime := time.Now().Add(time.Minute * 10)
	token, _ := helpers.GenerateJWT(user.ID, user.Role, expTime)

	var data = struct {
		Link string
	}{
		Link: "http://localhost:3000/recover_password?token=" + token,
	}

	err = mailService.SendMail("Alteracao de senha", requestBody.Email, "password_recover", data)

	if err != nil {
		log.Println(err)
	}

	c.Status(http.StatusCreated)
}

// ResetPassword
//
// @Summary      Reset user password
// @Tags         Users
// @Accept       json
// @Success      204
// @failure      400
// @failure      401
// @Param		 body	body	handlers.ResetPassword.request	true "body example"
// @Router		 /public/users/password [patch]
func ResetPassword(c *gin.Context) {
	type request struct {
		Token    string `json:"token"`
		Password string `json:"password" binding:"required,min=6,max=255"`
	}

	var requestBody = request{}
	accountsRepository := factories.GetAccountsRepository()

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

	claims, err := helpers.VerifyToken(requestBody.Token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	userId := claims.Id
	account, _ := accountsRepository.FindByUserID(userId)

	accountsRepository.SetPassword(account, helpers.GeneratePasswordHash(requestBody.Password))
}
