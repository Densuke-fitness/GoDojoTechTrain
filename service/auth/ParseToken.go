package auth

import (
	logger "github.com/sirupsen/logrus"

	"github.com/dgrijalva/jwt-go"
)

func DecodeToken(token string) (jwt.MapClaims, error) {

	t, err := parseToken(token)
	if err != nil {
		logger.Errorf("Error implementing parseToken: %s", err)
		return nil, err
	}
	return t.Claims.(jwt.MapClaims), nil
}

func parseToken(token string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		logger.Errorf("Error implementing jwt.Parse: %s", err)
		return t, err
	}
	return t, nil
}
