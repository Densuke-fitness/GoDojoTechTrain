// // //TODO:  not implemented
package gacha

import (
	"math/rand"

	"github.com/Densuke-fitness/GoDojoTechTrain/model/jwtUtil"
	logger "github.com/sirupsen/logrus"
)

type GachaResult struct {
	CharacterId int    `json:"characterID"`
	Name        string `json:"name"`
}

func DrawGacha(times int, token string) ([]GachaResult, error) {

	userId, err := jwtUtil.ExtractUserIdFromToken(token)
	if err != nil {
		logger.Errorf("Error ExtractUserIdFromToken: %s", err)
		return nil, err
	}

	LotteryRateMap, err := SelectLotteryRateAndCount()
	if err != nil {
		logger.Errorf("Error SelectLotteryRateAndCount: %s", err)
		return nil, err
	}

	var GachaResults []GachaResult

	for i := 1; i <= times; i++ {
		//the user draws randomly
		userRandNum := rand.Float64()
		rate := randChooseLotteryRate(userRandNum, LotteryRateMap)
		name, characterId, err := RandSelectCharacterByRate(rate)
		if err != nil {
			logger.Errorf("Error RandSelectCharacterByRate: %s", err)
			return nil, err
		}
		err = Insert(userId, characterId)
		if err != nil {
			logger.Errorf("Error Insert: %s", err)
			return nil, err
		}

		GachaResult := GachaResult{CharacterId: characterId, Name: name}
		GachaResults = append(GachaResults, GachaResult)
	}

	return GachaResults, err
}

func randChooseLotteryRate(userRandNum float64, LotteryRateMap map[float64]int) float64 {

	ApplicableboundaryVal := 0.0

	for rate, count := range LotteryRateMap {
		ApplicableboundaryVal += rate * float64(count)
		//When the rate is within the boundary value
		if userRandNum <= ApplicableboundaryVal {
			return rate
		}
	}
	//　If there is a problem with the selected value of user
	return 0.0
}
