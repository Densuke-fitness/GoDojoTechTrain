package gacha

import (
	"fmt"
	"testing"
)

func TestSelectLotteryRateList(t *testing.T) {

	//test :SelectLotteryRateList
	testLotteryRateList, err := SelectLotteryRateList()
	if err != nil {
		t.Errorf("Error implementing SelectLotteryRateList: %s", err.Error())
	}
	fmt.Println(testLotteryRateList)
}

func TestRandSelectCharacterByRate(t *testing.T) {

	//test :RandSelectCharacterByRate: 循環参照が起こるため値を設定
	tests := []struct {
		description string
		testRate    float64
		wantErr     bool
	}{
		//normal case
		{description: "Test Normal Case", testRate: 0.1, wantErr: false},
		//abnormal case
		{description: "Test Abnormal Case", testRate: 0.6, wantErr: true},
	}

	for id, tt := range tests {

		testCaseName := fmt.Sprintf("%v: %v", id+1, tt.description)

		t.Run(testCaseName, func(t *testing.T) {

			gachaResult, err := RandSelectCharacterByRate(tt.testRate)

			//test abnormal case
			if tt.wantErr {
				if err == nil {
					t.Errorf("Error implementing SelectLotteryRateList: %s", "Abnormal Case")
				}
			} else {

				//test normal case
				if err != nil {
					t.Errorf("Error implementing SelectLotteryRateList: %s", err.Error())
				}

				fmt.Println(gachaResult)
			}
		})
	}
}
