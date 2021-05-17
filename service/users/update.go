package users

import (
	"github.com/Densuke-fitness/GoDojoTechTrain/model"
	"github.com/Densuke-fitness/GoDojoTechTrain/repository/users"
	logger "github.com/sirupsen/logrus"
)

func UpdateUser(name string, userId int) error {

	userModelFromView := model.User{Name: name, Id: userId}

	//I used _ because I don't want to use the number of updates this time.
	_, err := users.UpdateNameById(userModelFromView)
	if err != nil {
		logger.Errorf("Error UpdateNameById: %s", err)
		return err
	}
	return nil
}
