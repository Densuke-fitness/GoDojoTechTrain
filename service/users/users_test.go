package users

import (
	"fmt"
	"testing"

	"github.com/Densuke-fitness/GoDojoTechTrain/service/tokenService/auth"
)

func TestUsers(t *testing.T) {

	//a userId
	tests := []struct {
		id   int
		name string
	}{
		{id: 1, name: "a"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("number:%v", tt.id)

		t.Run(testName, func(t *testing.T) {
			//test CreateUser
			userId, err := CreateUser(tt.name)
			if err != nil {
				t.Errorf("Error implementing CreateUser: %s", err.Error())
			}
			gotToken, err := auth.CreateToken(userId)
			if err != nil {
				t.Errorf("Error implementing auth.CreateToken: %s", err.Error())
			}

			//test GetUser
			userClaims, err := auth.DecodeToken(gotToken)
			if err != nil {
				t.Errorf("Error implementing auth.DecodeToken: %s", err.Error())
			}

			gotName, err := GetUser(userClaims.UserId)
			if err != nil {
				t.Errorf("Error implementing GetUser: %s", err.Error())
			}
			if gotName != tt.name {
				t.Errorf("Error implementing GetUser: %s", err.Error())
			}

			//test UpdateUser
			err = UpdateUser(gotName, userClaims.UserId)
			if err != nil {
				t.Errorf("Error implementing UpdateUser: %s", err.Error())
			}
		})

	}

}
