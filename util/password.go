package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns bcrypt hash string
func HashPassword(password string) (string, error) {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error to hashed the password %w ", err)
	}
	return string(hasedPassword), nil
}

// CheckPassword returns an error if current password is incorrect with the password in DB
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
