package users

import (
	"testing"
)

func TestRepository(t *testing.T) {
	t.Skip("TODO: not implemeted")

	// //a userId
	// tests := []struct {
	// 	userId  int
	// 	name    string
	// 	newName string
	// }{
	// 	{userId: 1, name: "a", newName: "new_a"},
	// 	{userId: 2, name: "ab", newName: "new_ab"},
	// 	{userId: 3, name: "AA", newName: "new_AA"},
	// 	{userId: 4, name: "あいう", newName: "new_あいう"},
	// 	{userId: 5, name: "123", newName: "new_123"},
	// 	{userId: 6, name: "１２３", newName: "new_１２３"},
	// }

	// for _, tt := range tests {
	// 	testName := fmt.Sprintf("number:%v", tt.userId)

	// 	t.Run(testName, func(t *testing.T) {
	// 		//test Insert
	// 		userModelFromView := model.User{Name: tt.name}
	// 		user, err := Insert(userModelFromView)
	// 		if err != nil {
	// 			t.Errorf("Error implementing Insert: %s", err.Error())
	// 		}

	// 		//test Select
	// 		gotUser, err := SelectNameById(*user)
	// 		if err != nil {
	// 			t.Errorf("Error implementing SelectNameById: %s", err.Error())
	// 		}
	// 		if gotUser.Name != tt.name {
	// 			t.Errorf(`Error SelectNameById: %v but want %q`, gotUser.Name, tt.name)
	// 		}

	// 		userModelFromView = model.User{Id: userModelFromView.Id, Name: tt.newName}

	// 		//test Update
	// 		_, err = UpdateNameById(userModelFromView)
	// 		if err != nil {
	// 			t.Errorf("Error implementing UpdateNameById: %s", err.Error())
	// 		}
	// 		gotUser, _ = SelectNameById(userModelFromView)
	// 		if gotUser.Name != tt.newName {
	// 			t.Errorf(`Error UpdateNameById(SelectNameById): %v but want %q`, gotUser.Name, tt.newName)
	// 		}
	// 	})

	// }

}
