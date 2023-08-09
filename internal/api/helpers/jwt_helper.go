package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(time.Minute * 60).Unix(),
	})
	tokenString, err := token.SignedString([]byte("secrectkey"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
