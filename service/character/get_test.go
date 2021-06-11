package character

import (
	"fmt"
	"testing"

	"github.com/Densuke-fitness/GoDojoTechTrain/service/gacha"
	"github.com/Densuke-fitness/GoDojoTechTrain/service/users"
)

func TestCharacter(t *testing.T) {

	//同値クラスの関係でテストケースを1つにしている
	tests := []struct {
		description  string
		testUserName string
		testTimes    int
	}{
		{description: "Test a series of API processes related to Character.", testUserName: "a"},
	}

	for id, tt := range tests {

		testCaseName := fmt.Sprintf("%v: %v", id+1, tt.description)

		t.Run(testCaseName, func(t *testing.T) {

			gotToken, _ := users.CreateUser(tt.testUserName)

			//gacha_testにて設定を10にしているためこちらも同様に設定
			_, _ = gacha.DrawGacha(tt.testTimes, gotToken)

			gotCharacters, err := GetCharacterList(gotToken)
			if err != nil {
				t.Errorf("Error GetCharacterList: %s", err.Error())
			}
			fmt.Println(gotCharacters)

		})
	}

}
