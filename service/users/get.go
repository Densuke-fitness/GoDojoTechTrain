package users

import (
	"github.com/Densuke-fitness/GoDojoTechTrain/model"
	"github.com/Densuke-fitness/GoDojoTechTrain/repository/users"
	"github.com/Densuke-fitness/GoDojoTechTrain/service/jwtUtil"
	logger "github.com/sirupsen/logrus"
)

func GetUser(token string) (string, error) {

	userId, err := jwtUtil.ExtractUserIdFromToken(token)
	if err != nil {
		logger.Errorf("Error ExtractUserIdFromToken: %s", err)
		return "", err
	}

	userModelFromView := model.User{Id: userId}

	//search name by using id
	user, err := users.SelectNameById(userModelFromView)
	if err != nil {
		logger.Errorf("Error SelectNameById: %s", err)
		return "", err
	}

	return user.Name, nil
}
