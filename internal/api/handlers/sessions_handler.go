package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gabrielmrts/mybooks-golang-api/internal/api/factories"
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/helpers"
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/models"
	"github.com/gin-gonic/gin"
)

type SessionRequestBody struct {
	Email    string
	Password string
}

// SessionStart
//
// @Summary      Start a session
// @Tags         Sessions
// @Accept       json
// @Produce      json
// @Success      201              {object}   handlers.SessionStart.response
// @failure      400
// @failure      401
// @Param		 body	body	SessionRequestBody	true "request example"
// @Router		 /public/sessions [post]
func SessionStart(c *gin.Context) {
	type response struct {
		Token string `json:"token"`
	}

	var requestBody SessionRequestBody
	usersRepository := factories.GetUsersRepository()
	accountsRepository := factories.GetAccountsRepository()
	emailVerificationRepository := factories.GetEmailVerificationRepository()
	mailService := factories.GetMailService()

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing email or password param"})
		return
	}

	passwordHash := helpers.GeneratePasswordHash(requestBody.Password)
	user, err := usersRepository.FindByEmailAndPassword(requestBody.Email, passwordHash)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	account, _ := accountsRepository.FindByUserID(user.ID)
	emailVerification, err := emailVerificationRepository.GetByEmail(account.Email)

	if err == nil {
		mailVariables := struct {
			Link string
			Name string
		}{
			Link: fmt.Sprintf("http://localhost:3000/verify?code=%s", emailVerification.Code),
			Name: user.Name,
		}

		if err := mailService.SendMail("Confirme seu email", account.Email, "email_confirmation", mailVariables); err != nil {
			log.Println("Error while sending email: ", err)
		}
	}

	if account.EmailVerified == false {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "email unconfirmed"})
		return
	}

	expTime := time.Now().Add(time.Hour * 1)
	token, _ := helpers.GenerateJWT(user.ID, user.Role, expTime)

	c.JSON(http.StatusCreated, gin.H{"token": token})
}

// SessionsMe
//
// @Summary      Get current session
// @Tags         Sessions
// @Produce      json
// @Success      200              {array}   handlers.SessionsMe.response
// @failure      400
// @failure      401
// @failure      404
// @Security Bearer
// @Router		 /private/sessions/me [get]
func SessionsMe(c *gin.Context) {
	type response struct {
		models.User    `json:"user"`
		models.Account `json:"account"`
	}

	usersRepository := factories.GetUsersRepository()
	userId := interface{}(c.MustGet("userId")).(uint)

	user, err := usersRepository.Get(userId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
