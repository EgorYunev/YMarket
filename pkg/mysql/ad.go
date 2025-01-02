package mysql

type Ad struct {
	id          int
	name        string
	description string
	owner       User
	price       int
}

func NewAd(id int, name string, description string, owner User, price int) *Ad {
	return &Ad{
		id:          id,
		name:        name,
		description: description,
		owner:       owner,
		price:       price,
	}
}
