package mysql

type User struct {
	id       int
	name     string
	password string
	ads      []Ad
}

func NewUser(id int, name string, password string) *User {
	return &User{
		id:       id,
		name:     name,
		password: password,
		ads:      []Ad{},
	}
}
