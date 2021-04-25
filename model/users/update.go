package users

import (
	"github.com/Densuke-fitness/GoDojoTechTrain/model/auth"
	logger "github.com/sirupsen/logrus"
)

func UpdateUser(name string, token string) error {

	decodedtoken, err := auth.DecodeToken(token)
	if err != nil {
		logger.Errorf("Error auth.DecodeToken: %s", err)
		return err
	}

	var userId int
	switch decodedtoken["user_id"] {
	case nil:
		logger.Warnf("Error decodedtoken['user_id']: %s", err)
		return err

	default:
		// extract userid
		userId = int(decodedtoken["user_id"].(float64))
	}

	//I used _ because I don't want to use the number of updates this time.
	_, err = UpdateNameById(name, userId)
	if err != nil {
		logger.Errorf("Error UpdateNameById: %s", err)
		return err
	}

	return nil
}
