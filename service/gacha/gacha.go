// // //TODO:  not implemented
package gacha

import (
	"math/rand"

	"github.com/Densuke-fitness/GoDojoTechTrain/model"
	"github.com/Densuke-fitness/GoDojoTechTrain/repository/character"
	"github.com/Densuke-fitness/GoDojoTechTrain/repository/gacha"
	logger "github.com/sirupsen/logrus"
)

func DrawGacha(times int, userId int) ([]model.Character, error) {

	userModelFromView := model.User{Id: userId}

	LotteryRateList, err := gacha.SelectLotteryRateList()
	if err != nil {
		logger.Errorf("Error SelectLotteryRateList: %s", err)
		return nil, err
	}

	var gachaResults []model.Character

	for i := 1; i <= times; i++ {
		//the user draws randomly
		userRandNum := rand.Float64()
		rate := RandChooseLotteryRate(userRandNum, LotteryRateList)
		gachaResult, err := gacha.RandSelectCharacterByRate(rate)
		if err != nil {
			logger.Errorf("Error RandSelectCharacterByRate: %s", err)
			return nil, err
		}

		maxSeq, err := character.SelectMaxSeqNum(userModelFromView, gachaResult)
		if err != nil {
			return nil, err
		}
		maxSeq += 1

		gachaResult.CharacterSeq = maxSeq

		err = character.Insert(userModelFromView, gachaResult)
		if err != nil {
			logger.Errorf("Error Insert: %s", err)
			return nil, err
		}

		gachaResults = append(gachaResults, gachaResult)
	}

	return gachaResults, err
}

func RandChooseLotteryRate(userRandNum float64, LotteryRateList []float64) float64 {

	applicableBoundaryVal := 0.0

	for _, rate := range LotteryRateList {
		applicableBoundaryVal += rate
		//When the rate is within the boundary value
		if userRandNum <= applicableBoundaryVal {
			return rate
		}
	}
	//　If there is a problem with the selected value of user
	logger.Error("Error randChooseLotteryRate")
	return 0.0
}
