package domain

type IAnimal interface {
	RegisterAnimal(name, animalType string) error
	DeleteAnimal(deleteId int) error
	GetAnimalInfo(id int) (domain.Animal, error)
	UpdateAnimal(updateId, newOwnerId int, newName, newType, newNickname string) error
	GetAnimals() ([]domain.Animal, error)
	GetNotAdoptedAnimals() ([]domain.Animal, error)
	GetOwnedAnimals(ownerId int) ([]domain.Animal, error)
	ChangeAnimalNickname(adoptId int, newNick string) error
}

type Animal struct {
	ID       int
	Name     string
	Type     string
	OwnerID  int
	Nickname string
}

var AnimalTypes []string
