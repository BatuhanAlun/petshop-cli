package service

import (
	"petshop/database"
	"petshop/domain"
)

func RegisterItem(iName string, iCost int) error {
	item := domain.Item{Name: iName, Cost: iCost}
	return database.RegisterItem(item)
}
