package auth

import (
	"github.com/dgrijalva/jwt-go"
)

func CheckToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return token, err
	}
	return token, nil
}
