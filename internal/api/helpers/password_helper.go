package helpers

import (
	"crypto/sha256"
	"fmt"
)

func GeneratePasswordHash(password string) string {
	passwordBytes := []byte(password)
	passwordHash := sha256.Sum256(passwordBytes)

	return fmt.Sprintf("%x", passwordHash)
}
