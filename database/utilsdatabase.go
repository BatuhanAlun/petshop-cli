package database

import (
	"os"
	"path/filepath"
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

					users.AddColumn(idCol)
					users.AddColumn(usernameCol)
					users.AddColumn(passwordCol)
					users.AddColumn(userRoleCol)
					// adding admin user
					hashedPass := pkg.CreateHash("admin")
					users.AddData([]string{"id", "username", "password", "role"}, []interface{}{0, "admin", hashedPass, "admin"})

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

		users.AddColumn(idCol)
		users.AddColumn(usernameCol)
		users.AddColumn(passwordCol)
		users.AddColumn(userRoleCol)
		// adding admin user
		hashedPass := pkg.CreateHash("admin")
		users.AddData([]string{"id", "username", "password", "role"}, []interface{}{0, "admin", hashedPass, "admin"})

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

		err := db.SaveDatabaseToFile()
		if err != nil {
			return false, err
		}
	}

	return true, nil
}
