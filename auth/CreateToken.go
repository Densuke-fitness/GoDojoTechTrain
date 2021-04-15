package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Tokenを生成する
// nameを渡せば返ってくる
func CreateToken(id int) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	token.Claims = jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	var secretKey = "secret"
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
