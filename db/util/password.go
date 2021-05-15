package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HassPassword(password string) (string, error) {
	passwordEncoded, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password %v", err)
	}
	return string(passwordEncoded), nil
}

func CheckPassword(password string, passwordEncoded string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordEncoded), []byte(password))
}
