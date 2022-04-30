package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/steveyiyo/simple-login/internal/tools"
)

var jwtKey = []byte(tools.RandomString(32))

type Token struct {
	jwt.StandardClaims
	UUID       string
	Permission string
}

// Generate Token
func GenerateToken(username string) (string, error) {
	expiresAt := time.Now().Add(24 * time.Hour).Unix()

	value := Token{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		UUID:       username,
		Permission: "1",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, value)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Validate Token
func ValidateToken(tokenString string) (bool, string) {
	var claims Token
	token, _ := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	var return_success bool
	var return_message string

	if !token.Valid {
		return_success = false
		return_message = "Token is invalid."
	} else {
		return_success = true
		return_message = "Token is valid."
	}

	return return_success, return_message
}
