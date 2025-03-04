package domain

type Animal struct {
	ID       int
	Name     string
	Type     string
	OwnerID  int
	Nickname string
}

var AnimalTypes []string
