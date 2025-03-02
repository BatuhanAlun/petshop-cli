package main

import (
	"fmt"
	"petshop/service"
)

func main() {
	for {
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. LogOut")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var username string
			var password string
			fmt.Println("Please type your Username:")
			fmt.Scanln(&username)
			fmt.Println("Please type your Password:")
			fmt.Scanln(&password)
			//userSessionId,
			_, userSessionRole, err := service.Login(username, password)
			if err != nil {
				fmt.Println(err)
				return
			}
			//admin loop
			if userSessionRole == "admin" {
				for {
					fmt.Println("1. Animal transactions")
					fmt.Println("2. Customer transactions")
					fmt.Println("3. Market transactions")
					var choice int
					fmt.Scanln(&choice)
					switch choice {
					case 1:
						//Animal Transactions
						fmt.Println("--==--==--Animal Transactions--==--==--")
						fmt.Println("1. Add New Animal")
						fmt.Println("2. Delete Animal")
						fmt.Println("3. Update Animal")
						fmt.Println("4. List  Animals")
						var choice int
						fmt.Scanln(&choice)
						switch choice {
						case 1:
							var name string
							var animalType string
							fmt.Println("--==--==--Add New Animal--==--==--")
							fmt.Println("Please Type Animal Name:")
							fmt.Scanln(&name)
							fmt.Println("Please type your Animals Type:")
							fmt.Scanln(&animalType)
							err := service.RegisterAnimal(name, animalType)
							if err != nil {
								fmt.Println(err)
							}

						case 2:
							var deleteId int
							fmt.Println("--==--==--Delete Animal--==--==--")
							fmt.Println("Please Type Animal id to Delete:")
							fmt.Scanln(&deleteId)
							err := service.DeleteAnimal(deleteId)
							if err != nil {
								fmt.Println(err)
							} else {
								fmt.Println("Succesfully deleted")
							}

						case 3:
							var updateId int
							fmt.Println("--==--==--Update Animal--==--==--")
							fmt.Println("Please Type Animal id to Update:")
							fmt.Scanln(&updateId)
							aniInfo, err := service.GetAnimalInfo(updateId)
							if err != nil {
								fmt.Println(err)
							}
							fmt.Println(aniInfo)
						case 4:
							fmt.Println("--==--==--List Animal--==--==--")
						}
					case 2:

					case 3:

					}
				}
			}

		case 2:
			fmt.Println("Kayıt işlemi")
			var username string
			var password string
			fmt.Println("Please type your Username:")
			fmt.Scanln(&username)
			fmt.Println("Please type your Password:")
			fmt.Scanln(&password)
			err := service.Register(username, password)
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			return
		}
	}
}
