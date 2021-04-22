package users

import (
	"fmt"
	"testing"
)

func TestRepository(t *testing.T) {

	//a userId
	tests := []struct {
		userId  int
		name    string
		newName string
	}{
		{userId: 1, name: "a", newName: "new_a"},
		{userId: 2, name: "ab", newName: "new_ab"},
		{userId: 3, name: "AA", newName: "new_AA"},
		{userId: 4, name: "あいう", newName: "new_あいう"},
		{userId: 5, name: "123", newName: "new_123"},
		{userId: 6, name: "１２３", newName: "new_１２３"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("number:%v", tt.userId)

		t.Run(testName, func(t *testing.T) {
			//test Insert
			id, err := Insert(tt.name)
			if err != nil {
				t.Errorf("Error implementing Insert: %s", err.Error())
			}

			//test Select
			gotName, err := SelectNameById(id)
			if err != nil {
				t.Errorf("Error implementing SelectNameById: %s", err.Error())
			}
			if gotName != tt.name {
				t.Errorf(`Error SelectNameById: %v but want %q`, gotName, tt.name)
			}

			//test Update
			_, err = UpdateNameById(tt.newName, id)
			if err != nil {
				t.Errorf("Error implementing UpdateNameById: %s", err.Error())
			}
			gotName, _ = SelectNameById(id)
			if gotName != tt.newName {
				t.Errorf(`Error UpdateNameById(SelectNameById): %v but want %q`, gotName, tt.newName)
			}
		})

	}

}
