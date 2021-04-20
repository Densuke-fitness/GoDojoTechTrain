package users

import (
	"github.com/Densuke-fitness/GoDojoTechTrain/model/auth"
	logger "github.com/sirupsen/logrus"
)

func CreateUser(name string) (string, error) {

	//database prosess
	id, err := Insert(name)

	if err != nil {
		logger.Errorf("Error Insert: %s", err)
		return "", err
	}

	//jwt prosess
	token, err := auth.CreateToken(id)
	if err != nil {
		logger.Errorf("Error auth.CreateToken: %s", err)
		return "", err
	}

	return token, nil
}
