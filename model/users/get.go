package users

import (
	"github.com/Densuke-fitness/GoDojoTechTrain/model/jwtUtil"
	logger "github.com/sirupsen/logrus"
)

func GetUser(token string) (string, error) {

	userId, err := jwtUtil.ExtractUserIdFromToken(token)
	if err != nil {
		logger.Errorf("Error ExtractUserIdFromToken: %s", err)
		return "", err
	}
	//search name by using id
	name, err := SelectNameById(userId)
	if err != nil {
		logger.Errorf("Error SelectNameById: %s", err)
		return "", err
	}

	return name, nil
}
