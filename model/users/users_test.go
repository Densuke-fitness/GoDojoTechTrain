package users

import (
	"fmt"
	"testing"
)

func TestUsers(t *testing.T) {

	//a userId
	tests := []struct {
		id   int
		name string
	}{
		{id: 1, name: "a"},
		{id: 2, name: "ab"},
		{id: 3, name: "AA"},
		{id: 4, name: "あいうう"},
		{id: 5, name: "123"},
		{id: 6, name: "１２３"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("number:%v", tt.id)

		t.Run(testName, func(t *testing.T) {
			//test CreateUser
			gotToken, err := CreateUser(tt.name)
			if err != nil {
				t.Errorf("Error implementing CreateUser: %s", err.Error())
			}

			//test GetUser
			gotName, err := GetUser(gotToken)
			if err != nil {
				t.Errorf("Error implementing GetUser: %s", err.Error())
			}
			if gotName != tt.name {
				t.Errorf("Error implementing GetUser: %s", err.Error())
			}

			//test UpdateUser
			err = UpdateUser(gotName, gotToken)
			if err != nil {
				t.Errorf("Error implementing UpdateUser: %s", err.Error())
			}
		})

	}

}
