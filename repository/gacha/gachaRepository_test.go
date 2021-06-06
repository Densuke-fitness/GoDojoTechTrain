package gacha

import (
	"fmt"
	"testing"
)

func TestSelectLotteryRateList(t *testing.T) {
	t.Skip()

	//test :SelectLotteryRateList
	testLotteryRateList, err := SelectLotteryRateList()
	if err != nil {
		t.Errorf("Error implementing SelectLotteryRateList: %s", err.Error())
	}
	fmt.Println(testLotteryRateList)
}

func TestRandSelectCharacterByRate(t *testing.T) {

	//test :RandSelectCharacterByRate: 循環参照が起こるため値を設定
	testRate := 0.1

	gachaResult, err := RandSelectCharacterByRate(testRate)
	if err != nil {
		t.Errorf("Error implementing SelectLotteryRateList: %s", err.Error())
	}
	fmt.Println(gachaResult)

}
