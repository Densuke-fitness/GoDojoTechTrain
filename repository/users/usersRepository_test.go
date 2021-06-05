package users

import (
	"fmt"
	"testing"

	"github.com/Densuke-fitness/GoDojoTechTrain/model"
)

func TestRepository(t *testing.T) {

	//a userId
	tests := []struct {
		id      int
		userId  int
		name    string
		newName string
	}{
		{userId: 1, name: "a", newName: "new_a"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("number:%v", tt.id)

		t.Run(testName, func(t *testing.T) {
			//test Insert: CreateUser
			testUserModelFromView := model.User{Name: tt.name}
			user, err := Insert(testUserModelFromView)
			if err != nil {
				t.Errorf("Error implementing Insert: %s", err.Error())
			}

			//test Select: GetUser
			user.Id = tt.id
			gotUser, err := SelectNameById(*user)
			if err != nil {
				t.Errorf("Error implementing SelectNameById: %s", err.Error())
			}
			if gotUser.Name != tt.name {
				t.Errorf(`Error SelectNameById: %v but want %q`, gotUser.Name, tt.name)
			}

			//test Update: UpdateUser
			testNewUserModelFromView := model.User{Id: tt.id, Name: tt.newName}
			_, err = UpdateNameById(testNewUserModelFromView)
			if err != nil {
				t.Errorf("Error implementing UpdateNameById: %s", err.Error())
			}
			gotUser, _ = SelectNameById(testNewUserModelFromView)
			if gotUser.Name != tt.newName {
				t.Errorf(`Error UpdateNameById(SelectNameById): %v but want %q`, gotUser.Name, tt.newName)
			}
		})

	}

}
