package auth

import (
	"fmt"
	"testing"
)

func TestCreateToken(t *testing.T) {

	//同値クラステストのため、1つのみのテストケース
	tests := []struct {
		userId int
	}{
		{userId: 1},
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
