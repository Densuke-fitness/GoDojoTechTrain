package users

import (
	"github.com/Densuke-fitness/GoDojoTechTrain/model"
	"github.com/Densuke-fitness/GoDojoTechTrain/repository/users"
	"github.com/Densuke-fitness/GoDojoTechTrain/service/jwtUtil"
	logger "github.com/sirupsen/logrus"
)

func UpdateUser(name string, token string) error {

	userId, err := jwtUtil.ExtractUserIdFromToken(token)
	if err != nil {
		logger.Errorf("Error ExtractUserIdFromToken: %s", err)
		return err
	}

	userModelFromView := model.User{Name: name, Id: userId}

	//I used _ because I don't want to use the number of updates this time.
	_, err = users.UpdateNameById(userModelFromView)
	if err != nil {
		logger.Errorf("Error UpdateNameById: %s", err)
		return err
	}

	return nil
}
