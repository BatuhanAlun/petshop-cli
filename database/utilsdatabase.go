package database

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"petshop/domain"
	"petshop/pkg"

	"github.com/BatuhanAlun/godb"
)

func FolderExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()

}

func FileExists(fileName string) bool {
	if filepath.Ext(fileName) != ".json" {
		fileName += ".json"
	}

	filePath := filepath.Join("DB", fileName)

	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
func checkTables(tableSlice []string) ([]string, bool) {
	returnValue := true
	faultyFiles := []string{}
	for _, v := range tableSlice {
		if !FileExists(v) {
			faultyFiles = append(faultyFiles, v)
		}
		returnValue = returnValue && FileExists(v)
	}
	return faultyFiles, returnValue
}

func Init() (bool, error) {
	mustKnownTables := []string{"users", "animals"}
	if FolderExists("DB") {
		missingTables, check := checkTables(mustKnownTables)
		if check {
			return true, nil
		} else {
			for _, v := range missingTables {
				switch v {
				case "users":
					db, err := godb.LoadDatabaseFromFile("DB")
					if err != nil {
						return false, err
					}
					users := godb.CreateTable("users")
					db.AddTable(users)

					idCol := godb.CreateColumn("id", "int", "PK")
					usernameCol := godb.CreateColumn("username", "string")
					passwordCol := godb.CreateColumn("password", "string")
					userRoleCol := godb.CreateColumn("role", "string")
					userMoneyCol := godb.CreateColumn("money", "int")

					users.AddColumn(idCol)
					users.AddColumn(usernameCol)
					users.AddColumn(passwordCol)
					users.AddColumn(userRoleCol)
					users.AddColumn(userMoneyCol)
					// adding admin user
					hashedPass := pkg.CreateHash("admin")
					users.AddData([]string{"id", "username", "password", "role", "money"}, []interface{}{0, "admin", hashedPass, "admin", 999999})

					err = db.SaveDatabaseToFile()
					if err != nil {
						return false, err
					}
				case "animals":
					db, err := godb.LoadDatabaseFromFile("DB")
					if err != nil {
						return false, err
					}
					animals := godb.CreateTable("animals")
					db.AddTable(animals)

					idCol := godb.CreateColumn("id", "int", "PK")
					animalNameCol := godb.CreateColumn("name", "string")
					animalTypeCol := godb.CreateColumn("type", "string")
					animalOwnerIdCol := godb.CreateColumn("ownerId", "string")
					animalNickCol := godb.CreateColumn("nickname", "string")

					animals.AddColumn(idCol)
					animals.AddColumn(animalNameCol)
					animals.AddColumn(animalTypeCol)
					animals.AddColumn(animalOwnerIdCol)
					animals.AddColumn(animalNickCol)

					err = db.SaveDatabaseToFile()
					if err != nil {
						return false, err
					}

				case "market":
					db, err := godb.LoadDatabaseFromFile("DB")
					if err != nil {
						return false, err
					}
					market := godb.CreateTable("market")
					db.AddTable(market)

					itemId := godb.CreateColumn("id", "int", "PK")
					itemName := godb.CreateColumn("name", "string")
					itemCost := godb.CreateColumn("cost", "int")

					market.AddColumn(itemId)
					market.AddColumn(itemName)
					market.AddColumn(itemCost)

					err = db.SaveDatabaseToFile()
					if err != nil {
						return false, err
					}
				}

			}

		}

	} else {
		db := godb.CreateDB("DB", "./")

		// INIT USERS TABLE
		users := godb.CreateTable("users")
		db.AddTable(users)

		idCol := godb.CreateColumn("id", "int", "PK")
		usernameCol := godb.CreateColumn("username", "string")
		passwordCol := godb.CreateColumn("password", "string")
		userRoleCol := godb.CreateColumn("role", "string")
		userMoneyCol := godb.CreateColumn("money", "int")

		users.AddColumn(idCol)
		users.AddColumn(usernameCol)
		users.AddColumn(passwordCol)
		users.AddColumn(userRoleCol)
		users.AddColumn(userMoneyCol)
		// adding admin user
		hashedPass := pkg.CreateHash("admin")
		users.AddData([]string{"id", "username", "password", "role", "money"}, []interface{}{0, "admin", hashedPass, "admin", 999999})

		// INIT ANIMALS TABLE
		animals := godb.CreateTable("animals")
		db.AddTable(animals)

		animalIdCol := godb.CreateColumn("id", "int", "PK")
		animalNameCol := godb.CreateColumn("name", "string")
		animalTypeCol := godb.CreateColumn("type", "string")
		animalOwnerIdCol := godb.CreateColumn("ownerId", "string")
		animalNickCol := godb.CreateColumn("nickname", "string")

		animals.AddColumn(animalIdCol)
		animals.AddColumn(animalNameCol)
		animals.AddColumn(animalTypeCol)
		animals.AddColumn(animalOwnerIdCol)
		animals.AddColumn(animalNickCol)

		// INIT MARKET TABLE
		market := godb.CreateTable("market")
		db.AddTable(market)

		itemId := godb.CreateColumn("id", "int", "PK")
		itemName := godb.CreateColumn("name", "string")
		itemCost := godb.CreateColumn("cost", "int")

		market.AddColumn(itemId)
		market.AddColumn(itemName)
		market.AddColumn(itemCost)

		err := db.SaveDatabaseToFile()
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func GetLastID(tablename string) int {
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
func IsIdExist(tablename string, id int) (bool, error) {
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

	for _, response := range responses {
		if response.Data.ID == id {
			return true, nil
		}
	}
	return false, nil
}
