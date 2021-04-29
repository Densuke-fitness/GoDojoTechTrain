package gacha

import (
	"fmt"
	"testing"

	"github.com/Densuke-fitness/GoDojoTechTrain/model/users"
)

func TestSelectLotteryRateAndCount(t *testing.T) {
	result, err := SelectLotteryRateAndCount()
	if err != nil {
		t.Errorf("Error SelectLotteryRateAndCount: %s", err.Error())
	}
	fmt.Println(result)

}
func TestRandSelectCharacterByRate(t *testing.T) {
	//Use the rate of the character that exists as test data.
	rate := 0.1

	name, characterId, err := RandSelectCharacterByRate(rate)
	if err != nil {
		t.Errorf("Error RandSelectCharacterByRate: %s", err.Error())
	}
	fmt.Printf(":%v, :%v", name, characterId)
}

func TestInsert(t *testing.T) {

	tests := []struct {
		id          int
		name        string
		characterId int
	}{
		{id: 1, characterId: 1},
		{id: 2, characterId: 2},
		{id: 3, characterId: 2},
	}

	testUser := "testUser"
	testUserId, _ := users.Insert(testUser)

	for _, tt := range tests {
		testName := fmt.Sprintf("number:%v", tt.id)

		t.Run(testName, func(t *testing.T) {
			//test Insert
			err := Insert(testUserId, tt.characterId)
			if err != nil {
				t.Errorf("Error implementing Insert: %s", err.Error())
			}
		})
	}
}
