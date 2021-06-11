package character

import (
	"fmt"
	"testing"

	"github.com/Densuke-fitness/GoDojoTechTrain/model"
	"github.com/Densuke-fitness/GoDojoTechTrain/repository/gacha"
	"github.com/Densuke-fitness/GoDojoTechTrain/repository/users"
)

func TestCharacterRepository(t *testing.T) {

	tests := []struct {
		description string
		testRate    float64
	}{
		{description: "Test the operation of the character repo process .", testRate: 0.1},
		{description: "Test if characterId is unique .", testRate: 0.1},
		{description: "Test to see if you can get a new kind of character.", testRate: 0.3},
	}

	testUserModelFromView := model.User{Name: "testUser"}
	//ユーザーの作成
	user, _ := users.Insert(testUserModelFromView)

	for id, tt := range tests {
		//ParallelTest
		tt := tt

		testCaseName := fmt.Sprintf("%v: %v", id+1, tt.description)

		t.Run(testCaseName, func(t *testing.T) {

			//ParallelTest
			t.Parallel()

			testGachaResult, _ := gacha.RandSelectCharacterByRate(tt.testRate)

			//test: SelectMaxSeqNum
			maxSeq, err := SelectMaxSeqNum(*user, testGachaResult)
			if err != nil {
				t.Errorf("Error implementing SelectMaxSeqNum: %s", err.Error())
			}

			//test: Insert
			maxSeq += 1
			testGachaResult.CharacterSeq = maxSeq

			err = Insert(*user, testGachaResult)
			if err != nil {
				t.Errorf("Error implementing Insert: %s", err.Error())
			}

			//test: SelectCharactersByUserId
			gotCharacters, err := SelectCharactersByUserId(*user)
			if err != nil {
				t.Errorf("Error implementing SelectCharactersByUserId: %s", err.Error())
			}
			fmt.Println(gotCharacters)

		})
	}
}
