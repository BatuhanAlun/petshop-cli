package database

import (
	"fmt"
	"petshop/domain"

	"github.com/BatuhanAlun/godb"
)

func RegisterItem(newItem domain.Item) error {

	var market *godb.Table
	newItemId := GetLastID("market.json")
	newItem.ID = newItemId
	db, err := godb.LoadDatabaseFromFile("DB")
	if err != nil {
		return err
	}
	for index, v := range db.Tables {
		if v.Name == "market" {
			market = db.Tables[index]
		}
	}

	err = market.AddData([]string{"id", "name", "cost"}, []interface{}{newItem.ID, newItem.Name, newItem.Cost})
	if err != nil {
		return err
	}
	err = db.SaveDatabaseToFile()
	if err != nil {
		return err
	}
	return nil

}

func DeleteItem(deleteId int) error {
	var market *godb.Table
	retval, err := IsIdExist("market.json", deleteId)
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
		if v.Name == "market" {
			market = db.Tables[index]
		}
	}

	err = market.Delete("id", deleteId)

	if err != nil {
		return err
	}
	err = db.SaveDatabaseToFile()
	if err != nil {
		return err
	}
	return nil
}

func GetItemInfo(id int) (domain.Item, error) {
	var itemInfo domain.Item
	_, err := IsIdExist("market.json", id)
	if err != nil {
		return itemInfo, err
	}
	db, err := godb.LoadDatabaseFromFile("DB")
	if err != nil {
		return itemInfo, err
	}
	//fetching item info in db
	for _, v := range db.Tables {
		if v.Name == "market" {
			for _, val := range v.Rows {
				if int(val.Data["id"].(float64)) == id {
					fmt.Println()
					itemInfo.ID = int(val.Data["id"].(float64))
					itemInfo.Name = val.Data["name"].(string)
					itemInfo.Cost = int(val.Data["cost"].(float64))
					return itemInfo, nil
				}
			}
		}
	}

	return itemInfo, fmt.Errorf("Item info cannot fetch")
}

func UpdateItem(updateId, newCost int, newName string) error {
	var updatedItem *godb.Table
	db, err := godb.LoadDatabaseFromFile("DB")
	if err != nil {
		return err
	}
	for index, v := range db.Tables {
		if v.Name == "market" {
			updatedItem = db.Tables[index]
		}
	}
	if newCost != 0 {
		err = updatedItem.Update("id", float64(updateId), "cost", float64(newCost))
		if err != nil {
			return err
		}
	}
	if newName != "" {
		updatedItem.Update("id", float64(updateId), "name", newName)
	}

	err = db.SaveDatabaseToFile()
	if err != nil {
		return err
	}
	return nil

}

func GetItemIdList() ([]int, error) {
	var idSlice []int
	db, err := godb.LoadDatabaseFromFile("DB")
	if err != nil {
		return nil, err
	}
	for _, v := range db.Tables {
		if v.Name == "market" {
			for _, val := range v.Rows {
				idSlice = append(idSlice, int(val.Data["id"].(float64)))
			}
		}
	}
	return idSlice, nil

}
