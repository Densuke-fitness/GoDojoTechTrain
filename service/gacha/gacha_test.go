package gacha

import (
	"fmt"
	"testing"

	"github.com/Densuke-fitness/GoDojoTechTrain/repository/gacha"
	"github.com/Densuke-fitness/GoDojoTechTrain/service/users"
)

func TestDrawGacha(t *testing.T) {

	tests := []struct {
		description  string
		testUserName string
		testTimes    int
	}{
		{description: "Test a series of API processes related to Users.", testUserName: "a"},
	}

	for id, tt := range tests {

		testCaseName := fmt.Sprintf("%v: %v", id+1, tt.description)

		t.Run(testCaseName, func(t *testing.T) {

			gotToken, _ := users.CreateUser(tt.testUserName)

			//テストケース: 同値クラステスト(数値は一旦10を採用)
			//TODO: 数値について議論するなら仕様を決める必要がある
			gachaResults, err := DrawGacha(tt.testTimes, gotToken)
			if err != nil {
				t.Errorf("Error DrawGacha: %s", err.Error())
			}
			fmt.Println(gachaResults)

		})
	}
}

func TestRandChooseLotteryRate(t *testing.T) {

	tests := []struct {
		description string
		userRandNum float64
	}{
		{description: "Test with small numbers", userRandNum: 0.1},
		{description: "Test with large numbers", userRandNum: 0.9},
	}

	//gchaRepositoryにてtest済み
	testLotteryRateList, _ := gacha.SelectLotteryRateList()

	for id, tt := range tests {

		testCaseName := fmt.Sprintf("%v: %v", id+1, tt.description)

		t.Run(testCaseName, func(t *testing.T) {

			rate := RandChooseLotteryRate(tt.userRandNum, testLotteryRateList)

			if rate <= 0 || 1 <= rate {
				t.Error("Error RandChooseLotteryRate")
			}
			fmt.Println(rate)

		})
	}

}
