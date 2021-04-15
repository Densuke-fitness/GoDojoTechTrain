package auth

import (
	"fmt"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

func TestCheckToken(*testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTg0ODc3MTIsImlkIjo3fQ.W81Obq7d4iTRSZcQWUPWCWkyJpkTeJ8mLA8rk5HI3Jk"
	t, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	decodedtoken := t.Claims.(jwt.MapClaims)
	fmt.Printf("%t", decodedtoken["id"])
}
