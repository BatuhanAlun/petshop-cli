package service

import (
	"petshop/database"
	"petshop/domain"
)

func RegisterAnimal(name, animalType string) error {

	animal := domain.Animal{Name: name, Type: animalType}
	return database.RegisterAnimal(animal)
}
func DeleteAnimal(deleteId int) error {
	return database.DeleteAnimal(deleteId)
}

func GetAnimalInfo(id int) (domain.Animal, error) {

	aniInfo, err := database.GetAnimalInfo(id)
	if err != nil {
		return aniInfo, err
	}
	return aniInfo, nil
}

func UpdateAnimal(updateId, newOwnerId int, newName, newType, newNickname string) error {
	err := database.UpdateAnimal(updateId, newOwnerId, newName, newType, newNickname)
	if err != nil {
		return err
	}
	return nil
}

func GetAnimals() ([]domain.Animal, error) {
	var aniInfoSlice []domain.Animal
	var tempInfo domain.Animal
	idList, err := database.GetAnimalIdList()
	if err != nil {
		return aniInfoSlice, err
	}
	for _, v := range idList {
		tempInfo, _ = GetAnimalInfo(v)
		aniInfoSlice = append(aniInfoSlice, tempInfo)
	}
	return aniInfoSlice, nil

}
