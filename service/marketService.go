package service

import (
	"petshop/database"
	"petshop/domain"
)

func RegisterItem(iName string, iCost int) error {
	item := domain.Item{Name: iName, Cost: iCost}
	return database.RegisterItem(item)
}

func DeleteItem(deleteId int) error {
	return database.DeleteItem(deleteId)
}

func GetItemInfo(id int) (domain.Item, error) {

	aniInfo, err := database.GetItemInfo(id)
	if err != nil {
		return aniInfo, err
	}
	return aniInfo, nil
}

func UpdateItem(updateId, newCost int, newName string) error {

	err := database.UpdateItem(updateId, newCost, newName)
	if err != nil {
		return err
	}
	return nil

}

func GetItems() ([]domain.Item, error) {
	var itemInfoSlice []domain.Item
	var tempInfo domain.Item
	idList, err := database.GetItemIdList()
	if err != nil {
		return itemInfoSlice, err
	}
	for _, v := range idList {
		tempInfo, _ = GetItemInfo(v)
		itemInfoSlice = append(itemInfoSlice, tempInfo)
	}
	return itemInfoSlice, nil

}

func GetRecords() ([]domain.Records, error) {
	var recordInfoSlice []domain.Records
	var tempInfo domain.Records
	idList, err := database.GetRecordsIdList()
	if err != nil {
		return recordInfoSlice, err
	}
	for _, v := range idList {
		tempInfo, _ = database.GetRecordInfo(v)
		recordInfoSlice = append(recordInfoSlice, tempInfo)
	}
	return recordInfoSlice, nil
}
