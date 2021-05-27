package auth

import (
	"fmt"

	"github.com/Densuke-fitness/GoDojoTechTrain/service/tokenService"
	logger "github.com/sirupsen/logrus"

	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	UserId int
}

func DecodeToken(token string) (*UserClaims, error) {

	t, err := parseToken(token)
	if err != nil {
		logger.Errorf("Error implementing parseToken: %s", err)
		return nil, err
	}

	mapClaims := t.Claims.(jwt.MapClaims)

	userClaims := convertToUserClaims(mapClaims)
	if userClaims == nil {
		err = fmt.Errorf("%s", "including invalid fields")
		logger.Errorf("Error implementing convertToUserClaims: %s", err)
		return nil, err
	}

	return userClaims, nil
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

func convertToUserClaims(mapClaims jwt.MapClaims) *UserClaims {

	var userClaims UserClaims

	//userId
	if mapClaims[tokenService.USER_ID] == nil {
		logger.Warnf("Error  mapClaims['%s']: %s", tokenService.USER_ID, mapClaims[tokenService.USER_ID])
		return nil
	}
	userId := int(mapClaims[tokenService.USER_ID].(float64))
	userClaims.UserId = userId

	return &userClaims
}
