package auth

import (
	"log"

	"github.com/dgrijalva/jwt-go"
)

func parseToken(token string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		log.Fatalf("Error implementing jwt.Parse: %s", err)
		return t, err
	}
	return t, nil
}

func DecodeToken(token string) (jwt.MapClaims, error) {

	t, err := parseToken(token)
	if err != nil {
		log.Fatalf("Error implementing parseToken: %s", err)
		return nil, err
	}
	return t.Claims.(jwt.MapClaims), nil
}
