package gacha

import (
	"fmt"
	"testing"

	"github.com/Densuke-fitness/GoDojoTechTrain/model/users"
)

func TestGacha(t *testing.T) {

	//test CreateUser
	const testUser = "testUser"
	gotToken, _ := users.CreateUser(testUser)

	const testTimes = 10

	GachaResults, err := DrawGacha(testTimes, gotToken)
	if err != nil {
		t.Errorf("Error DrawGacha: %s", err.Error())
	}
	fmt.Println(GachaResults)

}
