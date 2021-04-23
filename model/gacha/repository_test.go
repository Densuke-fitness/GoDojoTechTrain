package gacha

import (
	"fmt"
	"testing"
)

func TestRepository(t *testing.T) {

	//a id
	tests := []struct {
		id             int
		userLotteryNum float64
	}{
		{id: 1, userLotteryNum: 0.1},
		{id: 2, userLotteryNum: 0.2},
		{id: 3, userLotteryNum: 0.4},
		{id: 4, userLotteryNum: 0.6},
		{id: 5, userLotteryNum: 0.8},
		{id: 6, userLotteryNum: 1.0},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("number:%v", tt.id)

		t.Run(testName, func(t *testing.T) {

			//test Select
			gotName, gotId, err := Select(tt.userLotteryNum)
			if err != nil {
				t.Errorf("Error implementing SelectCharacterIdDrawn: %s", err.Error())
			}
			fmt.Println(gotName, gotId)
		})
	}
}
