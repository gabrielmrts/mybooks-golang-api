package handlers

import (
	"errors"
	"net/http"

	"github.com/gabrielmrts/mybooks-golang-api/internal/api/factories"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VerifyEmailRequestBody struct {
	Code string `json:"code"`
}

// VerifyEmail
//
// @Summary      Verify a user email
// @Tags         Users
// @Accept       json
// @Success      204
// @failure      400
// @failure      401
// @Param		 body	body	models.EmailVerification	true "body example"
// @Router		 /public/users/email [patch]
func VerifyEmail(c *gin.Context) {
	var requestBody VerifyEmailRequestBody
	emailVerificationRepository := factories.GetEmailVerificationRepository()
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

	emailVerify, err := emailVerificationRepository.GetByCode(requestBody.Code)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "invalid code"})
		return
	}

	account, _ := accountsRepository.FindByEmail(emailVerify.Email)
	accountsRepository.SetEmailVerified(account, emailVerify.Email)
	c.Status(http.StatusNoContent)
}
