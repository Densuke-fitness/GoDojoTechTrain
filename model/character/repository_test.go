package character

import (
	"fmt"
	"testing"

	"github.com/Densuke-fitness/GoDojoTechTrain/model/gacha"
	"github.com/Densuke-fitness/GoDojoTechTrain/model/users"
)

func TestSelectCharactersById(t *testing.T) {

	testUser := "testUser"
	testUserId, _ := users.Insert(testUser)

	testCharacterId := 1
	_ = gacha.Insert(testUserId, testCharacterId)

	gotCharacters, err := SelectCharactersById(testUserId)
	if err != nil {
		t.Errorf("Error SelectCharactersById: %s", err.Error())
	}
	fmt.Println(gotCharacters)
}
