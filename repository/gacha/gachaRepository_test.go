package gacha

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/Densuke-fitness/GoDojoTechTrain/service/gacha"
)

func TestGachaRepository(t *testing.T) {

	//test :SelectLotteryRateList
	testLotteryRateList, err := SelectLotteryRateList()
	if err != nil {
		t.Errorf("Error implementing SelectLotteryRateList: %s", err.Error())
	}
	fmt.Println(testLotteryRateList)

	//test :RandSelectCharacterByRate
	testUserRandNum := rand.Float64()
	testRate := gacha.RandChooseLotteryRate(testUserRandNum, testLotteryRateList)

	gachaResult, err := RandSelectCharacterByRate(testRate)
	if err != nil {
		t.Errorf("Error implementing SelectLotteryRateList: %s", err.Error())
	}
	fmt.Println(gachaResult)

}
