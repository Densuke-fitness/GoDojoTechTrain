package character

import (
	"github.com/Densuke-fitness/GoDojoTechTrain/model"
	"github.com/Densuke-fitness/GoDojoTechTrain/repository/character"
	logger "github.com/sirupsen/logrus"
)

func GetCharacterList(userId int) ([]model.Character, error) {

	userModel := model.User{Id: userId}

	characters, err := character.SelectCharactersByUserId(userModel)
	if err != nil {
		logger.Errorf("Error SelectCharactersById: %s", err)
		return nil, err
	}

	return characters, nil
}
