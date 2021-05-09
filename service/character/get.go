package character

import (
	"github.com/Densuke-fitness/GoDojoTechTrain/model"
	"github.com/Densuke-fitness/GoDojoTechTrain/repository/character"
	"github.com/Densuke-fitness/GoDojoTechTrain/service/jwtUtil"
	logger "github.com/sirupsen/logrus"
)

func GetCharacterList(token string) ([]model.Character, error) {
	userId, err := jwtUtil.ExtractUserIdFromToken(token)
	if err != nil {
		logger.Errorf("Error ExtractUserIdFromToken: %s", err)
		return nil, err
	}

	userModelFromView := model.User{Id: userId}

	characters, err := character.SelectCharactersByUserId(userModelFromView)
	if err != nil {
		logger.Errorf("Error SelectCharactersById: %s", err)
		return nil, err
	}

	return characters, nil
}
