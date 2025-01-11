package models

type User struct {
	Id       int
	Name     string
	Password string
	Ads      []Ad
}

func NewUser(name, password string) User {
	return User{
		Name:     name,
		Password: password,
		Ads:      make([]Ad, 0),
	}
}
