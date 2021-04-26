package users

import (
	"github.com/Densuke-fitness/GoDojoTechTrain/model/jwtUtil"
	logger "github.com/sirupsen/logrus"
)

func UpdateUser(name string, token string) error {

	userId, err := jwtUtil.ExtractUserIdFromToken(token)
	if err != nil {
		logger.Errorf("Error ExtractUserIdFromToken: %s", err)
		return err
	}

	//I used _ because I don't want to use the number of updates this time.
	_, err = UpdateNameById(name, userId)
	if err != nil {
		logger.Errorf("Error UpdateNameById: %s", err)
		return err
	}

	return nil
}
