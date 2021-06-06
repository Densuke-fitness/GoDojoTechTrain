package character

import (
	"fmt"
	"testing"

	"github.com/Densuke-fitness/GoDojoTechTrain/service/gacha"
	"github.com/Densuke-fitness/GoDojoTechTrain/service/users"
)

func TestCharacter(t *testing.T) {

	const testUser = "testUser"
	gotToken, _ := users.CreateUser(testUser)

	const testTimes = 10
	_, _ = gacha.DrawGacha(testTimes, gotToken)

	gotCharacters, err := GetCharacterList(gotToken)
	if err != nil {
		t.Errorf("Error GetCharacterList: %s", err.Error())
	}
	fmt.Println(gotCharacters)

}
