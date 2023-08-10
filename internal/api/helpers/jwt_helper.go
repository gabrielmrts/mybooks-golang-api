package helpers

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userId uint, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Id":   userId,
		"Role": role,
		"Exp":  time.Now().Add(time.Minute * 60).Unix(),
	})
	tokenString, err := token.SignedString([]byte("secrectkey"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

type JWTClaims struct {
	Id   uint   `json:"id"`
	Exp  int64  `json:"exp"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func VerifyToken(signedToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secrectkey"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return nil, errors.New("malformed token")
	} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		return nil, errors.New("invalid signature")
	} else if errors.Is(err, jwt.ErrTokenExpired) {
		return nil, errors.New("expired token")
	} else {
		return nil, errors.New("invalid token")
	}
}
