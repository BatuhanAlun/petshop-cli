package database

import (
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
