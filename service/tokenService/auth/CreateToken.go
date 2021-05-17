package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	logger "github.com/sirupsen/logrus"
)

//Pass a unique id and return a token
func CreateToken(userId int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}

	var secretKey = "secret"
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		logger.Errorf("Error auth.CreateToken: %s", err)
		return "", err
	}
	return tokenString, nil
}
