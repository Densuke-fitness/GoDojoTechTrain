package auth

import (
	"fmt"
	"testing"
)

func TestCreateToken(t *testing.T) {

	//a id
	tests := []struct {
		id   int
		want string
	}{
		{id: 1},
		{id: 2},
		{id: 3},
		{id: 4},
		{id: 5},
		{id: 6},
		{id: 7},
		{id: 8},
		{id: 9},
		{id: 10},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("number:%v", tt.id)

		t.Run(name, func(t *testing.T) {
			_, err := CreateToken(tt.id)
			if err != nil {
				t.Errorf("Error implementing CreateToken: %s", err.Error())
			}

		})

	}

}
