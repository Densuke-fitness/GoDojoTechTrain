package tokenService

import (
	"github.com/Densuke-fitness/GoDojoTechTrain/service/tokenService/auth"
	logger "github.com/sirupsen/logrus"
)

func ExtractUserIdFromToken(token string) (int, error) {
	decodedtoken, err := auth.DecodeToken(token)
	if err != nil {
		logger.Errorf("Error auth.DecodeToken: %s", err)
		// Used -1 , not 0 beacuse of using -1 in model/auth/repository.go
		return -1, err
	}

	if decodedtoken[USER_ID] == nil {
		logger.Warnf("Error decodedtoken['%s']: %s", USER_ID, err)
		return -1, err
	}

	// extract userid
	userId := int(decodedtoken[USER_ID].(float64))

	return userId, nil
}
