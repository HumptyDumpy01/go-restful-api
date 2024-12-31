package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "534aacfceb99dab96590bfd49a4fecf81b892e2a4a6b8f0af1aae5a62b44a5b2"

// GenerateToken before using jwt, install it, if did not: go get -u github.com/golang-jwt/jwt/v5
func GenerateToken(email string, userId int64) (string, error) {
	fmt.Println(`Printing userId:`, userId)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	// create a const and gen a secret key by e.g. running "openssl rand -hex 32"
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, errors.New("couldn't parse the token")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	fmt.Println(`Executing claims:`, claims)

	if !ok {
		return 0, errors.New("invalid token claim")
	}

	//email := claims["email"].(string)
	userIdFloat, ok := claims["userId"].(float64)
	fmt.Println("userIdFloat executing:", userIdFloat)

	if !ok {
		return 0, errors.New("userId claim is not a float64")
	}

	userId := int64(userIdFloat)

	return userId, nil
}
