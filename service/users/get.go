package users

import (
	"github.com/Densuke-fitness/GoDojoTechTrain/model"
	"github.com/Densuke-fitness/GoDojoTechTrain/repository/users"
	logger "github.com/sirupsen/logrus"
)

func GetUser(userId int) (string, error) {

	userModelFromView := model.User{Id: userId}

	//search name by using id
	user, err := users.SelectNameById(userModelFromView)
	if err != nil {
		logger.Errorf("Error SelectNameById: %s", err)
		return "", err
	}

	return user.Name, nil
}
