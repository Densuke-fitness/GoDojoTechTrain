package gacha

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/Densuke-fitness/GoDojoTechTrain/repository/gacha"
	"github.com/Densuke-fitness/GoDojoTechTrain/service/users"
)

func TestDrawGacha(t *testing.T) {

	//test CreateUser
	const testUser = "testUser"
	gotToken, _ := users.CreateUser(testUser)

	//テストケース: 同値クラステスト(数値は一旦10を採用)
	//TODO: 数値について議論するなら仕様を決める必要がある
	const testTimes = 10

	GachaResults, err := DrawGacha(testTimes, gotToken)
	if err != nil {
		t.Errorf("Error DrawGacha: %s", err.Error())
	}
	fmt.Println(GachaResults)
}

func TestRandChooseLotteryRate(t *testing.T) {

	//gchaRepositoryにてtest済み
	testLotteryRateList, _ := gacha.SelectLotteryRateList()
	testUserRandNum := rand.Float64()

	rate := RandChooseLotteryRate(testUserRandNum, testLotteryRateList)

	if rate <= 0 || 1 <= rate {
		t.Error("Error RandChooseLotteryRate")
	}
}
