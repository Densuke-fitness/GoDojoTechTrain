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
		id          int
		characterId int
	}{
		{id: 1, characterId: 1},
		{id: 2, characterId: 1},
		{id: 3, characterId: 2},
	}

	testUserModelFromView := model.User{Name: "testUser"}
	//ユーザーの作成
	user, _ := users.Insert(testUserModelFromView)

	for _, tt := range tests {
		//ParallelTest
		tt := tt

		testName := fmt.Sprintf("number:%v", tt.id)

		t.Run(testName, func(t *testing.T) {

			//ParallelTest
			t.Parallel()

			testRate := 0.1
			testGachaResult, _ := gacha.RandSelectCharacterByRate(testRate)

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
