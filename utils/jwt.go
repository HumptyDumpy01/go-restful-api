package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "534aacfceb99dab96590bfd49a4fecf81b892e2a4a6b8f0af1aae5a62b44a5b2"

// GenerateToken before using jwt, install it, if did not: go get -u github.com/golang-jwt/jwt/v5
func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	// create a const and gen a secret key by e.g. running "openssl rand -hex 32"
	return token.SignedString(secretKey)
}
