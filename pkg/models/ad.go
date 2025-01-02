package models

type Ad struct {
	Id          int
	Name        string
	Description string
	Owner       User
	Price       int
	UserId      int
}
