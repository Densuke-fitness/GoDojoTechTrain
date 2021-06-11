package users

import (
	"fmt"
	"testing"

	"github.com/Densuke-fitness/GoDojoTechTrain/model"
)

func TestRepository(t *testing.T) {

	//user_test.goでテストケースを1つにしているためこちらも同様に設定
	tests := []struct {
		description     string
		testUserId      int
		testUserName    string
		testNewUserName string
	}{
		{description: "Test a series of API processes related to Users.",
			testUserId: 1, testUserName: "a", testNewUserName: "new_a"},
	}

	for id, tt := range tests {
		testCaseName := fmt.Sprintf("%v: %v", id+1, tt.description)

		t.Run(testCaseName, func(t *testing.T) {
			//test Insert: CreateUser
			testUserModelFromView := model.User{Name: tt.testUserName}
			user, err := Insert(testUserModelFromView)
			if err != nil {
				t.Errorf("Error implementing Insert: %s", err.Error())
			}

			//test Select: GetUser
			user.Id = tt.testUserId
			gotUser, err := SelectNameById(*user)
			if err != nil {
				t.Errorf("Error implementing SelectNameById: %s", err.Error())
			}
			if gotUser.Name != tt.testUserName {
				t.Errorf(`Error SelectNameById: %v but want %q`, gotUser.Name, tt.testUserName)
			}

			//test Update: UpdateUser
			testNewUserModelFromView := model.User{Id: tt.testUserId, Name: tt.testNewUserName}
			_, err = UpdateNameById(testNewUserModelFromView)
			if err != nil {
				t.Errorf("Error implementing UpdateNameById: %s", err.Error())
			}
			gotUser, _ = SelectNameById(testNewUserModelFromView)
			if gotUser.Name != tt.testNewUserName {
				t.Errorf(`Error UpdateNameById(SelectNameById): %v but want %q`, gotUser.Name, tt.testNewUserName)
			}
		})

	}

}
