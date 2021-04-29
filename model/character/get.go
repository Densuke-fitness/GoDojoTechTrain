package character

import (
	"github.com/Densuke-fitness/GoDojoTechTrain/model/jwtUtil"
	logger "github.com/sirupsen/logrus"
)

func GetCharacterList(token string) ([]Character, error) {
	userId, err := jwtUtil.ExtractUserIdFromToken(token)
	if err != nil {
		logger.Errorf("Error ExtractUserIdFromToken: %s", err)
		return nil, err
	}

	Characters, err := SelectCharactersById(userId)
	if err != nil {
		logger.Errorf("Error SelectCharactersById: %s", err)
		return nil, err
	}

	return Characters, nil
}
