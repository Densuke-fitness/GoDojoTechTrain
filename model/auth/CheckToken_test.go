package auth

import (
	"fmt"
	"testing"
)

func TestCheckToken(t *testing.T) {

	//a id
	tests := []struct {
		n  int
		id int
	}{
		{n: 1, id: 1},
		{n: 2, id: 2},
		{n: 3, id: 3},
		{n: 4, id: 4},
		{n: 5, id: 5},
		{n: 6, id: 6},
		{n: 7, id: 7},
		{n: 8, id: 8},
		{n: 9, id: 9},
		{n: 10, id: 10},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("number:%v", tt.n)

		t.Run(name, func(t *testing.T) {
			got, _ := CreateToken(tt.id)
			_, err := checkToken(got)
			if err != nil {
				t.Errorf("Error implementing checkToken: %s", err.Error())
			}
		})
	}

}

func TestDecodeToken(t *testing.T) {

	//a id
	tests := []struct {
		n    int
		id   int
		want int
	}{
		{n: 1, id: 1, want: 1},
		{n: 2, id: 2, want: 2},
		{n: 3, id: 3, want: 3},
		{n: 4, id: 4, want: 4},
		{n: 5, id: 5, want: 5},
		{n: 6, id: 6, want: 6},
		{n: 7, id: 7, want: 7},
		{n: 8, id: 8, want: 8},
		{n: 9, id: 9, want: 9},
		{n: 10, id: 10, want: 10},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("number:%v", tt.n)

		t.Run(name, func(t *testing.T) {
			token, _ := CreateToken(tt.id)
			got, err := DecodeToken(token)
			if err != nil {
				t.Errorf("Error implementing checkToken: %s", err.Error())
			}
			got_id := int(got["id"].(float64))
			if got_id != tt.want {
				t.Errorf(`Error DecodeToken: %v but want %q`, got_id, tt.want)
			}
		})
	}

}
