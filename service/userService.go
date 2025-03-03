package service

import (
	"fmt"
	"petshop/database"
	"petshop/domain"
	"petshop/pkg"
)

func Register(username, password string) error {
	hashedPass := pkg.CreateHash(password)
	user := domain.User{Username: username, Password: hashedPass, Role: "customer", Money: 0}
	return database.SaveUser(user)
}

func Login(username, password string) (int, string, error) {
	userInfo, err := database.IsUserExist(username)
	var errInt int = -1
	var errString string = "err"
	if err != nil {
		return errInt, errString, err
	}

	retval := pkg.ComparePassword(password, userInfo.Password)
	if retval != true {
		return errInt, errString, fmt.Errorf("wrong password")
	}

	return userInfo.ID, userInfo.Role, err
}

// same
func DeleteCustomer(deleteId int) error {
	return database.DeleteCustomer(deleteId)
}

// same
func GetUserInfo(id int) (domain.User, error) {

	aniInfo, err := database.FetchUserInfoById(id)
	if err != nil {
		return aniInfo, err
	}
	return aniInfo, nil
}

func UpdateUser(updateId, newMoney int, newUsername, newPassword string) error {
	var hashedPass string
	if newPassword != "" {
		hashedPass = pkg.CreateHash(newPassword)
	} else {
		hashedPass = ""
	}

	err := database.UpdateUser(updateId, newMoney, newUsername, hashedPass)
	if err != nil {
		return err
	}
	return nil

}

func GetUsers() ([]domain.User, error) {
	var usrInfoSlice []domain.User
	var tempInfo domain.User
	idList, err := database.GetUserIdList()
	if err != nil {
		return usrInfoSlice, err
	}
	for _, v := range idList {
		tempInfo, _ = GetUserInfo(v)
		usrInfoSlice = append(usrInfoSlice, tempInfo)
	}
	return usrInfoSlice, nil

}

func AdoptAnimal(id, userId int) error {
	err := database.AdoptAnimal(id, userId)
	return err
}

func BuyItem(itemId, userId int) error {
	userInfo, err := GetUserInfo(userId)
	if err != nil {
		return err
	}
	itemInfo, err := GetItemInfo(itemId)
	if err != nil {
		return err
	}
	if userInfo.Money < itemInfo.Cost {
		return fmt.Errorf("User Do not have Enough Money!")
	}
	newMoney := userInfo.Money - itemInfo.Cost
	err = UpdateUser(userId, newMoney, "", "")
	if err != nil {
		fmt.Println("buraya giriyo")
		return err
	}
	err = database.BuyItem(itemId, userId)
	if err != nil {
		return err
	}
	return nil
}

func AddMoney(userId, money int) error {
	return database.AddMoney(userId, money)
}
