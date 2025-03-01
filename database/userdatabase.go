package database

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
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
	id := getLastID("users.json")
	erer := users.AddData([]string{"id", "username", "password", "role"}, []interface{}{id, user.Username, user.Password, user.Role})
	if erer != nil {
		fmt.Println(erer)
	}
	rrr := db.SaveDatabaseToFile()
	if rrr != nil {
		fmt.Println(rrr)
	}

	return nil
}

func getLastID(tablename string) int {
	filePath := filepath.Join("DB", tablename)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var responses []domain.Response
	if err := json.Unmarshal(bytes, &responses); err != nil {
		log.Fatal(err)
	}

	var lastID int
	for _, response := range responses {
		if response.Data.ID > lastID {
			lastID = response.Data.ID
		}
	}

	return lastID + 1
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
					return userInfo, nil
				}
			}
		}
	}
	return emptyUser, fmt.Errorf("user not found")
}
