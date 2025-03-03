package database

import (
	"fmt"
	"petshop/domain"

	"github.com/BatuhanAlun/godb"
)

// func Init() bool {
// 	if FolderExists("DB") {
// 		return true
// 	} else {
// 		db := godb.CreateDB("DB", "./")
// 		users := godb.CreateTable("users")
// 		db.AddTable(users)

// 		idCol := godb.CreateColumn("id", "int", "PK")
// 		usernameCol := godb.CreateColumn("username", "string")
// 		passwordCol := godb.CreateColumn("password", "string")
// 		userRoleCol := godb.CreateColumn("role", "string")

// 		users.AddColumn(idCol)
// 		users.AddColumn(usernameCol)
// 		users.AddColumn(passwordCol)
// 		users.AddColumn(userRoleCol)
// 		// adding admin user
// 		hashedPass := pkg.CreateHash("admin")
// 		users.AddData([]string{"id", "username", "password", "role"}, []interface{}{0, "admin", hashedPass, "admin"})

// 		err := db.SaveDatabaseToFile()
// 		if err != nil {
// 			return false
// 		}
// 		return true
// 	}
// }

func SaveUser(user domain.User) error {
	var users *godb.Table

	check, _ := Init()
	if !check {
		return fmt.Errorf("DB Cannot Created!")
	}
	db, err := godb.LoadDatabaseFromFile("DB")
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("DB Cannot Pulled")
	}
	for index, v := range db.Tables {
		if v.Name == "users" {
			users = db.Tables[index]
		}
	}
	id := GetLastID("users.json")
	erer := users.AddData([]string{"id", "username", "password", "role", "money"}, []interface{}{id, user.Username, user.Password, user.Role, user.Money})
	if erer != nil {
		fmt.Println(erer)
	}
	rrr := db.SaveDatabaseToFile()
	if rrr != nil {
		fmt.Println(rrr)
	}

	return nil
}

func IsUserExist(username string) (domain.User, error) {
	var emptyUser domain.User
	var userInfo domain.User
	check, _ := Init()
	if !check {
		return emptyUser, fmt.Errorf("DB Cannot Created!")
	}
	db, err := godb.LoadDatabaseFromFile("DB")
	if err != nil {
		fmt.Println(err)
		return emptyUser, fmt.Errorf("DB Cannot Pulled")
	}
	//fetching user info in db
	for _, v := range db.Tables {
		if v.Name == "users" {
			for _, val := range v.Rows {
				if val.Data["username"] == username {
					userInfo.ID = int(val.Data["id"].(float64))
					userInfo.Username = val.Data["username"].(string)
					userInfo.Password = val.Data["password"].(string)
					userInfo.Role = val.Data["role"].(string)
					userInfo.Money = int(val.Data["money"].(float64))
					return userInfo, nil
				}
			}
		}
	}
	return emptyUser, fmt.Errorf("user not found")
}

func DeleteCustomer(deleteId int) error {
	var users *godb.Table
	retval, err := IsIdExist("users.json", deleteId)
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
		if v.Name == "users" {
			users = db.Tables[index]
		}
	}

	err = users.Delete("id", deleteId)

	if err != nil {
		return err
	}
	err = db.SaveDatabaseToFile()
	if err != nil {
		return err
	}
	return nil
}

func FetchUserInfoById(id int) (domain.User, error) {
	var emptyUser domain.User
	var userInfo domain.User
	check, _ := Init()
	if !check {
		return emptyUser, fmt.Errorf("DB Cannot Created!")
	}
	db, err := godb.LoadDatabaseFromFile("DB")
	if err != nil {
		fmt.Println(err)
		return emptyUser, fmt.Errorf("DB Cannot Pulled")
	}
	//fetching user info in db
	for _, v := range db.Tables {
		if v.Name == "users" {
			for _, val := range v.Rows {
				if int(val.Data["id"].(float64)) == id {
					userInfo.ID = int(val.Data["id"].(float64))
					userInfo.Username = val.Data["username"].(string)
					userInfo.Password = val.Data["password"].(string)
					userInfo.Role = val.Data["role"].(string)
					userInfo.Money = int(val.Data["money"].(float64))
					return userInfo, nil
				}
			}
		}
	}
	return emptyUser, fmt.Errorf("user not found")
}

func UpdateUser(updateId, newMoney int, newUsername, newPassword string) error {
	var updatedUser *godb.Table
	db, err := godb.LoadDatabaseFromFile("DB")
	if err != nil {
		return err
	}
	for index, v := range db.Tables {
		if v.Name == "users" {
			updatedUser = db.Tables[index]
		}
	}
	if newMoney != 0 {
		err = updatedUser.Update("id", float64(updateId), "money", float64(newMoney))
		if err != nil {
			return err
		}
	}
	if newUsername != "" {
		updatedUser.Update("id", float64(updateId), "username", newUsername)
	}
	if newPassword != "" {
		updatedUser.Update("id", float64(updateId), "password", newPassword)
	}

	err = db.SaveDatabaseToFile()
	if err != nil {
		return err
	}
	return nil

}

func GetUserIdList() ([]int, error) {
	var idSlice []int
	db, err := godb.LoadDatabaseFromFile("DB")
	if err != nil {
		return nil, err
	}
	for _, v := range db.Tables {
		if v.Name == "users" {
			for _, val := range v.Rows {
				idSlice = append(idSlice, int(val.Data["id"].(float64)))
			}
		}
	}
	return idSlice, nil

}

func BuyItem(itemId, userId int) error {
	var marketRecord *godb.Table
	db, err := godb.LoadDatabaseFromFile("DB")
	if err != nil {
		return err
	}
	for _, v := range db.Tables {
		if v.Name == "marketRecords" {
			marketRecord = v
		}
	}
	lastId := GetLastID("marketRecords.json")
	err = marketRecord.AddData([]string{"id", "ownerId", "itemId"}, []interface{}{lastId, userId, itemId})
	if err != nil {
		return err
	}
	err = db.SaveDatabaseToFile()
	if err != nil {
		return err
	}
	return nil
}
