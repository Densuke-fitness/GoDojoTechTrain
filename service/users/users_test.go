package users

import (
	"fmt"
	"testing"

	"github.com/Densuke-fitness/GoDojoTechTrain/service/tokenService/auth"
)

func TestUsers(t *testing.T) {

	//同値クラスの関係でテストケースを1つにしている:異常系はcontrallerでバリデートしている
	tests := []struct {
		description string
		userName    string
	}{
		{
			description: "Test a series of API processes related to Users.",
			userName:    "a",
		},
	}

	for id, tt := range tests {
		testCaseName := fmt.Sprintf("%v: %v", id+1, tt.description)

		t.Run(testCaseName, func(t *testing.T) {
			//test CreateUser
			userId, err := CreateUser(tt.userName)
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

			gotUserName, err := GetUser(userClaims.UserId)
			if err != nil {
				t.Errorf("Error implementing GetUser: %s", err.Error())
			}
			if gotUserName != tt.userName {
				t.Errorf("Error implementing GetUser: %s", err.Error())
			}

			//test UpdateUser
			err = UpdateUser(gotUserName, userClaims.UserId)
			if err != nil {
				t.Errorf("Error implementing UpdateUser: %s", err.Error())
			}
		})

	}

}
