package users

import (
	"github.com/Densuke-fitness/GoDojoTechTrain/model"
	"github.com/Densuke-fitness/GoDojoTechTrain/repository/users"
	"github.com/Densuke-fitness/GoDojoTechTrain/service/auth"
	logger "github.com/sirupsen/logrus"
)

func CreateUser(name string) (string, error) {
	userModelFromView := model.User{Name: name}
	//database prosess

	user, err := users.Insert(userModelFromView)

	if err != nil {
		logger.Errorf("Error Insert: %s", err)
		return "", err
	}

	//jwt prosess
	token, err := auth.CreateToken(user.Id)
	if err != nil {
		logger.Errorf("Error auth.CreateToken: %s", err)
		return "", err
	}

	return token, nil
}
