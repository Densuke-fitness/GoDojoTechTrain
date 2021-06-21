package users

import (
	"fmt"
	"testing"

	"github.com/Densuke-fitness/GoDojoTechTrain/model"
)

func TestRepository(t *testing.T) {

	//user_test.goでテストケースを1つにしているためこちらも同様に設定
	tests := []struct {
		description string
		userId      int
		userName    string
		newUserName string
	}{
		{description: "Test a series of API processes related to Users.",
			userId: 1, userName: "a", newUserName: "new_a"},
	}

	for id, tt := range tests {
		testCaseName := fmt.Sprintf("%v: %v", id+1, tt.description)

		t.Run(testCaseName, func(t *testing.T) {
			//test Insert: CreateUser
			testUserModelFromView := model.User{Name: tt.userName}
			user, err := Insert(testUserModelFromView)
			if err != nil {
				t.Errorf("Error implementing Insert: %s", err.Error())
			}

			//test Select: GetUser
			gotUser, err := SelectNameById(*user)
			if err != nil {
				t.Errorf("Error implementing SelectNameById: %s", err.Error())
			}
			if gotUser.Name != tt.userName {
				t.Errorf(`Error SelectNameById: %v but want %q`, gotUser.Name, tt.userName)
			}

			//test Update: UpdateUser
			testNewUserModelFromView := model.User{Id: tt.userId, Name: tt.newUserName}
			_, err = UpdateNameById(testNewUserModelFromView)
			if err != nil {
				t.Errorf("Error implementing UpdateNameById: %s", err.Error())
			}
			gotUser, _ = SelectNameById(testNewUserModelFromView)
			if gotUser.Name != tt.newUserName {
				t.Errorf(`Error UpdateNameById(SelectNameById): %v but want %q`, gotUser.Name, tt.newUserName)
			}
		})

	}

}
