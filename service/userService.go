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
