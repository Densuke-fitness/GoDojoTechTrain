package users

import (
	"github.com/Densuke-fitness/GoDojoTechTrain/model"
	"github.com/Densuke-fitness/GoDojoTechTrain/repository/users"
	logger "github.com/sirupsen/logrus"
)

func CreateUser(name string) (int, error) {
	userModelFromView := model.User{Name: name}
	//database prosess
	user, err := users.Insert(userModelFromView)

	if err != nil {
		logger.Errorf("Error Insert: %s", err)
		return 0, err
	}

	return user.Id, nil
}
