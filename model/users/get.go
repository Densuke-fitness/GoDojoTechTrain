package users

import (
	"github.com/Densuke-fitness/GoDojoTechTrain/model/auth"
	logger "github.com/sirupsen/logrus"
)

func GetUser(token string) (string, error) {

	decodedtoken, err := auth.DecodeToken(token)
	if err != nil {
		logger.Errorf("Error auth.DecodeToken: %s", err)
		return "", err
	}

	var userId int
	switch decodedtoken["user_id"] {
	case nil:
		logger.Warnf("Error decodedtoken['user_id']: %s", err)
		return "", err
	default:
		// extract userid
		userId = int(decodedtoken["user_id"].(float64))
	}

	//search name by using id
	name, err := SelectNameById(userId)
	if err != nil {
		logger.Errorf("Error SelectNameById: %s", err)
		return "", err
	}

	return name, nil
}
