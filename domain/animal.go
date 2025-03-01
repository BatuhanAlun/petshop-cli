package domain

type IAnimal interface {
	GetID() int
	GetName() string
	SetName(newName string)
	GetType() string
	SetType(animalType string)
	GetOwnerID() int
	SetOwnerID(ownerID int)
}

type Animal struct {
	ID       int
	Name     string
	Type     string
	OwnerID  int
	Nickname string
}
