package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// In a real application, you might want to handle this error differently
		return ""
	}
	return string(hashedBytes)
}
