package models

type User struct {
	Id       int
	Name     string
	Password string
	Ads      []Ad
}
