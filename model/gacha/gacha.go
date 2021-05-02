// // //TODO:  not implemented
package gacha

import (
	"math/rand"
	"strconv"

	"github.com/Densuke-fitness/GoDojoTechTrain/model/jwtUtil"
	logger "github.com/sirupsen/logrus"
)

type GachaResult struct {
	CharacterId string `json:"characterID"`
	Name        string `json:"name"`
}

func DrawGacha(times int, token string) ([]GachaResult, error) {

	userId, err := jwtUtil.ExtractUserIdFromToken(token)
	if err != nil {
		logger.Errorf("Error ExtractUserIdFromToken: %s", err)
		return nil, err
	}

	LotteryRateList, err := SelectLotteryRateAndCount()
	if err != nil {
		logger.Errorf("Error SelectLotteryRateAndCount: %s", err)
		return nil, err
	}

	var gachaResults []GachaResult

	for i := 1; i <= times; i++ {
		//the user draws randomly
		userRandNum := rand.Float64()
		rate := randChooseLotteryRate(userRandNum, LotteryRateList)
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

		characterIdStr := strconv.Itoa(characterId)

		gachaResult := GachaResult{CharacterId: characterIdStr, Name: name}
		gachaResults = append(gachaResults, gachaResult)
	}

	return gachaResults, err
}

func randChooseLotteryRate(userRandNum float64, LotteryRateList []float64) float64 {

	ApplicableboundaryVal := 0.0

	for _, rate := range LotteryRateList {
		ApplicableboundaryVal += rate
		//When the rate is within the boundary value
		if userRandNum <= ApplicableboundaryVal {
			return rate
		}
	}
	//　If there is a problem with the selected value of user
	logger.Error("Error randChooseLotteryRate")
	return 0.0
}
