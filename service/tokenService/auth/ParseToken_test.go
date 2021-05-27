package auth

import (
	"testing"
)

func TestAuth(t *testing.T) {
	t.Skip("TODO: not implemented")
	// tests := []struct {
	// 	n      int
	// 	userId int
	// 	want   int
	// }{
	// 	{n: 1, userId: 1, want: 1},
	// 	{n: 2, userId: 2, want: 2},
	// 	{n: 3, userId: 3, want: 3},
	// 	{n: 4, userId: 4, want: 4},
	// 	{n: 5, userId: 5, want: 5},
	// 	{n: 6, userId: 6, want: 6},
	// 	{n: 7, userId: 7, want: 7},
	// 	{n: 8, userId: 8, want: 8},
	// 	{n: 9, userId: 9, want: 9},
	// 	{n: 10, userId: 10, want: 10},
	// }

	// for _, tt := range tests {
	// 	name := fmt.Sprintf("number:%v", tt.n)

	// 	t.Run(name, func(t *testing.T) {
	// 		token, err := CreateToken(tt.userId)
	// 		if err != nil {
	// 			t.Errorf("Error implementing parseToken: %s", err.Error())
	// 		}
	// 		got, err := DecodeToken(token)
	// 		if err != nil {
	// 			t.Errorf("Error implementing parseToken: %s", err.Error())
	// 		}
	// 		gotUserId := int(got["user_id"].(float64))
	// 		if gotUserId != tt.want {
	// 			t.Errorf(`Error DecodeToken: %v but want %q`, gotUserId, tt.want)
	// 		}
	// 	})
	// }

}
