package auth

import (
	"fmt"
	"testing"
)

func TestAuth(t *testing.T) {

	//同値クラステストのため、1つのみのテストケース
	tests := []struct {
		n      int
		userId int
		want   int
	}{
		{n: 1, userId: 1, want: 1},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("number:%v", tt.n)

		t.Run(name, func(t *testing.T) {
			token, err := CreateToken(tt.userId)
			if err != nil {
				t.Errorf("Error implementing CreateToken: %s", err.Error())
			}
			got, err := DecodeToken(token)
			if err != nil {
				t.Errorf("Error implementing DecodeToken: %s", err.Error())
			}
			gotUserId := int(got["user_id"].(float64))
			if gotUserId != tt.want {
				t.Errorf(`Error DecodeToken: %v but want %q`, gotUserId, tt.want)
			}
		})
	}

}
