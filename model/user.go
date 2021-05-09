package model

type User struct {
	Id   int
	Name string
}

func (u *User) Clone(id int) *User {

	user := User{
		Id:   id,
		Name: u.Name,
	}

	return &user
}
