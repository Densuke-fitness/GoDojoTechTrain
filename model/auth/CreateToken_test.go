package auth

import (
	"fmt"
	"testing"
)

func TestCreateToken(t *testing.T) {

	//a userId
	tests := []struct {
		userId int
	}{
		{userId: 1},
		{userId: 2},
		{userId: 3},
		{userId: 4},
		{userId: 5},
		{userId: 6},
		{userId: 7},
		{userId: 8},
		{userId: 9},
		{userId: 10},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("number:%v", tt.userId)

		t.Run(name, func(t *testing.T) {
			_, err := CreateToken(tt.userId)
			if err != nil {
				t.Errorf("Error implementing CreateToken: %s", err.Error())
			}

		})

	}

}
