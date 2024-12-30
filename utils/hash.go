package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}

	return string(fromPassword), err
}