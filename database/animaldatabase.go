package database

import (
	"fmt"
	"petshop/domain"

	"github.com/BatuhanAlun/godb"
)

func RegisterAnimal(newAnimal domain.Animal) error {
	var animals *godb.Table
	animalid := GetLastID("animals.json")
	newAnimal.ID = animalid
	db, err := godb.LoadDatabaseFromFile("DB")
	if err != nil {
		return err
	}
	for index, v := range db.Tables {
		if v.Name == "animals" {
			animals = db.Tables[index]
		}
	}
	newAnimal.OwnerID = -1
	newAnimal.Nickname = "noNickNameGiven"
	err = animals.AddData([]string{"id", "name", "type", "ownerId", "nickname"}, []interface{}{newAnimal.ID, newAnimal.Name, newAnimal.Type, newAnimal.OwnerID, newAnimal.Nickname})
	if err != nil {
		return err
	}
	err = db.SaveDatabaseToFile()
	if err != nil {
		return err
	}
	return nil
}

func DeleteAnimal(deleteId int) error {
	var animals *godb.Table
	retval, err := IsIdExist("animals.json", deleteId)
	if err != nil {
		return err
	}
	if !retval {
		return fmt.Errorf("id not exist")
	}

	db, err := godb.LoadDatabaseFromFile("DB")

	if err != nil {
		return err
	}

	for index, v := range db.Tables {
		if v.Name == "animals" {
			animals = db.Tables[index]
		}
	}

	err = animals.Delete("id", deleteId)

	if err != nil {
		return err
	}
	err = db.SaveDatabaseToFile()
	if err != nil {
		return err
	}
	return nil
}

func GetAnimalInfo(id int) (domain.Animal, error) {
	var animalInfo domain.Animal
	_, err := IsIdExist("animals.json", id)
	if err != nil {
		return animalInfo, err
	}
	db, err := godb.LoadDatabaseFromFile("DB")
	if err != nil {
		return animalInfo, err
	}
	//fetching user info in db
	for _, v := range db.Tables {
		if v.Name == "animals" {
			for _, val := range v.Rows {
				if int(val.Data["id"].(float64)) == id {
					fmt.Println()
					animalInfo.ID = int(val.Data["id"].(float64))
					animalInfo.Name = val.Data["name"].(string)
					animalInfo.Nickname = val.Data["nickname"].(string)
					animalInfo.OwnerID = int(val.Data["ownerId"].(float64))
					animalInfo.Type = val.Data["type"].(string)
					return animalInfo, nil
				}
			}
		}
	}

	return animalInfo, fmt.Errorf("animal info cannot fetch")
}
