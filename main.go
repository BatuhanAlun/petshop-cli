package main

import (
	"fmt"
	"log"
	"petshop/service"
)

func main() {
	for {
		fmt.Println("1. Login")
		fmt.Println("2. Register")

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

			userSessionId, userSessionRole, err := service.Login(username, password)
			if err != nil {
				fmt.Println(err)
				return
			}
			service.AddLog(userSessionId, userSessionRole, "Logged In")
			//admin loop
			if userSessionRole == "admin" {
				for {
					fmt.Println("1. Animal transactions")
					fmt.Println("2. Customer transactions")
					fmt.Println("3. Market transactions")
					fmt.Println("4. Log Out")
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
						fmt.Println("5. Return Upper-Menu")
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
							service.AddLog(userSessionId, userSessionRole, "Added Animal")

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
							service.AddLog(userSessionId, userSessionRole, "Deleted Animal")
						case 3:
							var updateId, newOwnerId int
							var newName, newType, newNickname string
							fmt.Println("--==--==--Update Animal--==--==--")
							fmt.Println("Please Type Animal id to Update:")
							fmt.Scanln(&updateId)
							aniInfo, err := service.GetAnimalInfo(updateId)
							if err != nil {
								fmt.Println(err)
							}
							fmt.Println(aniInfo)
							fmt.Println("Please insert the new values if you don't want to change the value just press Enter!")
							fmt.Println("Please Type Animal Name to Update:")
							fmt.Scanln(&newName)
							fmt.Println("Please Type Animal Type to Update:")
							fmt.Scanln(&newType)
							fmt.Println("Please Type Animal Owner id to Update:")
							_, err = fmt.Scanln(&newOwnerId)
							if err != nil {
								newOwnerId = -1
							}
							fmt.Println("Please Type Animal Nickname to Update:")
							fmt.Scanln(&newNickname)
							err = service.UpdateAnimal(updateId, newOwnerId, newName, newType, newNickname)
							if err != nil {
								fmt.Println(err)
							} else {
								fmt.Println("Successufly updated!")
								animal, err := service.GetAnimalInfo(updateId)
								if err != nil {
									fmt.Println(err)
								}
								fmt.Println(animal)
							}
							service.AddLog(userSessionId, userSessionRole, "Updated Animal")
						case 4:
							fmt.Println("--==--==--List Animal--==--==--")
							animalList, err := service.GetAnimals()
							if err != nil {
								log.Fatal(err)
							}
							for index, _ := range animalList {
								fmt.Printf("%d. Animal Info : %v\n", index, animalList[index])
								fmt.Println("-------------------------------")
							}
							service.AddLog(userSessionId, userSessionRole, "Listed Animals")
						case 5:
							break
						}
					case 2:
						//User Transactions
						fmt.Println("--==--==--Customer Transactions--==--==--")
						fmt.Println("1. Add New Customer")
						fmt.Println("2. Delete Customer")
						fmt.Println("3. Update Customer")
						fmt.Println("4. List  Customer")
						fmt.Println("5. Return Upper-Menu")
						var choice int
						fmt.Scanln(&choice)
						switch choice {
						case 1:
							var username string
							var password string
							fmt.Println("--==--==--Add New Customer--==--==--")
							fmt.Println("Please Type Customer Username:")
							fmt.Scanln(&username)
							fmt.Println("Please type your Customer Password:")
							fmt.Scanln(&password)
							err := service.Register(username, password)
							if err != nil {
								fmt.Println(err)
							}
							service.AddLog(userSessionId, userSessionRole, "Added New Customer")
						case 2:
							var deleteId int
							fmt.Println("--==--==--Delete Customer--==--==--")
							fmt.Println("Please Type Customer id to Delete:")
							fmt.Scanln(&deleteId)
							err := service.DeleteCustomer(deleteId)
							if err != nil {
								fmt.Println(err)
							} else {
								fmt.Println("Succesfully deleted")
							}
							service.AddLog(userSessionId, userSessionRole, "Deleted Customer")

						case 3:
							var updateId, newMoney int
							var newUsername, newPassword string
							fmt.Println("--==--==--Update Customer--==--==--")
							fmt.Println("Please Type Customer id to Update:")
							fmt.Scanln(&updateId)
							usrInfo, err := service.GetUserInfo(updateId)
							if err != nil {
								fmt.Println(err)
							}
							fmt.Println(usrInfo)
							fmt.Println("Please insert the new values if you don't want to change the value just press Enter!")
							fmt.Println("Please Type Username to Update:")
							fmt.Scanln(&newUsername)
							fmt.Println("Please Type User Password to Update:")
							fmt.Scanln(&newPassword)
							fmt.Println("Please Type User Money to Update:")
							_, err = fmt.Scanln(&newMoney)
							if err != nil {
								newMoney = 0
							}
							err = service.UpdateUser(updateId, newMoney, newUsername, newPassword)
							if err != nil {
								fmt.Println(err)
							} else {
								fmt.Println("Successufly updated!")
								animal, err := service.GetUserInfo(updateId)
								if err != nil {
									fmt.Println(err)
								}
								fmt.Println(animal)
							}
							service.AddLog(userSessionId, userSessionRole, "Updated Customer")
						case 4:
							fmt.Println("--==--==--List Users--==--==--")
							userList, err := service.GetUsers()
							if err != nil {
								log.Fatal(err)
							}
							for index, _ := range userList {
								fmt.Printf("%d. User Info : %v\n", index, userList[index])
								fmt.Println("-------------------------------")
							}
							service.AddLog(userSessionId, userSessionRole, "Listed Users")
						case 5:
							break

						}
					case 3:
						//Market Transactions
						fmt.Println("--==--==--Market Transactions--==--==--")
						fmt.Println("1. Add New Item")
						fmt.Println("2. Delete Item")
						fmt.Println("3. Update Item")
						fmt.Println("4. List  Item")
						fmt.Println("5. Return Upper-Menu")
						var choice int
						fmt.Scanln(&choice)
						switch choice {
						case 1:
							var iName string
							var iCost int
							fmt.Println("--==--==--Add New Item--==--==--")
							fmt.Println("Please Type Item Name:")
							fmt.Scanln(&iName)
							fmt.Println("Please type your Item Cost:")
							fmt.Scanln(&iCost)
							err := service.RegisterItem(iName, iCost)
							if err != nil {
								fmt.Println(err)
							}
							service.AddLog(userSessionId, userSessionRole, "Added New Item")

						case 2:
							var deleteId int
							fmt.Println("--==--==--Delete Item--==--==--")
							fmt.Println("Please Type Item id to Delete:")
							fmt.Scanln(&deleteId)
							err := service.DeleteItem(deleteId)
							if err != nil {
								fmt.Println(err)
							} else {
								fmt.Println("Succesfully deleted")
							}
							service.AddLog(userSessionId, userSessionRole, "Deleted Item")

						case 3:
							var updateId, newCost int
							var newName string
							fmt.Println("--==--==--Update Item--==--==--")
							fmt.Println("Please Type Item id to Update:")
							fmt.Scanln(&updateId)
							itemInfo, err := service.GetItemInfo(updateId)
							if err != nil {
								fmt.Println(err)
							}
							fmt.Println(itemInfo)
							fmt.Println("Please insert the new values if you don't want to change the value just press Enter!")
							fmt.Println("Please Type New Item Name to Update:")
							fmt.Scanln(&newName)
							fmt.Println("Please Type New Cost of the Item to Update:")
							_, err = fmt.Scanln(&newCost)
							if err != nil {
								newCost = 0
							}
							err = service.UpdateItem(updateId, newCost, newName)
							if err != nil {
								fmt.Println(err)
							} else {
								fmt.Println("Successufly updated!")
								animal, err := service.GetItemInfo(updateId)
								if err != nil {
									fmt.Println(err)
								}
								fmt.Println(animal)
							}
							service.AddLog(userSessionId, userSessionRole, "Updated Item")

						case 4:
							fmt.Println("--==--==--List Items--==--==--")
							itemList, err := service.GetItems()
							if err != nil {
								log.Fatal(err)
							}
							for index, _ := range itemList {
								fmt.Printf("%d. Item Info : %v\n", index, itemList[index])
								fmt.Println("-------------------------------")
							}
							service.AddLog(userSessionId, userSessionRole, "Listed Items")
						case 5:
							break
						}
					case 4:
						return
					}
				}
			} else if userSessionRole == "customer" {
				for {
					fmt.Println("--==--==--Customer Page--==--==--")
					fmt.Println("1. Animal Page")
					fmt.Println("2. Market Page")
					fmt.Println("3. Add Money Page")
					fmt.Println("4. Log Out")
					var choice int
					fmt.Scanln(&choice)
					switch choice {
					case 1:
						fmt.Println("--==--==--Animal Page--==--==--")
						fmt.Println("1. Adopt an Animal")
						fmt.Println("2. Give Nickname")
						fmt.Println("3. List Your Animals")
						fmt.Println("4. List  All Animals")
						fmt.Println("5. Return Upper Menu")
						var choice int
						fmt.Scanln(&choice)
						switch choice {
						case 1:
							var adoptId int
							fmt.Println("--==--==--Adopt Animal--==--==--")
							animalList, err := service.GetNotAdoptedAnimals()
							if err != nil {
								log.Fatal(err)
							}
							if animalList == nil {
								fmt.Println("All animals Adopted!")
								break
							}
							for index, _ := range animalList {
								fmt.Printf("%d. Animal Info : %v\n", index, animalList[index])
								fmt.Println("-------------------------------")
							}
							fmt.Println("Please Type Animal ID You Want to Adopt")
							fmt.Scanln(&adoptId)
							err = service.AdoptAnimal(adoptId, userSessionId)
							if err != nil {
								fmt.Println(err)
							}
							fmt.Println("Succesfull")
							service.AddLog(userSessionId, userSessionRole, "Adopted Animal")
						case 2:
							var adoptId int
							var newNick string
							fmt.Println("--==--==--Give Nick to Animal--==--==--")
							animalList, err := service.GetOwnedAnimals(userSessionId)
							if err != nil {
								log.Fatal(err)
							}
							if animalList == nil {
								fmt.Println("You do not adopted any Animals!!")
								break
							}
							for index, _ := range animalList {
								fmt.Printf("%d. Animal Info : %v\n", index, animalList[index])
								fmt.Println("-------------------------------")
							}
							fmt.Println("Please Type Animal ID You Want to Change Nickname")
							fmt.Scanln(&adoptId)
							fmt.Println("Please Type Animal Nick you Want to Set")
							fmt.Scanln(&newNick)
							err = service.ChangeAnimalNickname(adoptId, newNick)
							if err != nil {
								fmt.Println(err)
							}
							fmt.Println("Succesfull")
							service.AddLog(userSessionId, userSessionRole, "Gived Nickname to Animal")

						case 3:
							fmt.Println("--==--==--Your Animals--==--==--")
							animalList, err := service.GetOwnedAnimals(userSessionId)
							if err != nil {
								log.Fatal(err)
							}
							if animalList == nil {
								fmt.Println("You do not adopted any Animals!!")

							}
							for index, _ := range animalList {
								fmt.Printf("%d. Animal Info : %v\n", index, animalList[index])
								fmt.Println("-------------------------------")
							}
							service.AddLog(userSessionId, userSessionRole, "Listed Owned Animals")
						case 4:
							fmt.Println("--==--==--All Animals--==--==--")
							animalList, err := service.GetAnimals()
							if err != nil {
								log.Fatal(err)
							}
							for index, _ := range animalList {
								fmt.Printf("%d. Animal Info : %v\n", index, animalList[index])
								fmt.Println("-------------------------------")
							}
							service.AddLog(userSessionId, userSessionRole, "Listed All Animals")

						case 5:
							break
						}

					case 2:
						fmt.Println("--==--==--Market Page--==--==--")
						fmt.Println("1. Buy Item")
						fmt.Println("2. Check Inventory")
						fmt.Println("3. List Items")
						fmt.Println("4. Return Upper-Menu")
						var choice int
						fmt.Scanln(&choice)
						switch choice {

						case 1:
							var buyItemId int
							fmt.Println("--==--==--Buy Menu--==--==--")
							itemList, err := service.GetItems()
							if err != nil {
								log.Fatal(err)
							}
							for index, _ := range itemList {
								fmt.Printf("%d. Item Info : %v\n", index, itemList[index])
								fmt.Println("-------------------------------")
							}
							fmt.Println("Please Type Item id to Buy")
							fmt.Scanln(&buyItemId)
							err = service.BuyItem(buyItemId, userSessionId)
							if err != nil {
								fmt.Println(err)
							}
							service.AddLog(userSessionId, userSessionRole, "Buyed Item")
						case 2:
							fmt.Println("--==--==--Check Inventory--==--==--")
							var recordPrintList []string
							itemList, err := service.GetItems()
							if err != nil {
								log.Fatal(err)
							}
							userRecords, err := service.GetRecords()
							if err != nil {
								log.Fatal(err)
							}
							for _, v := range userRecords {
								if v.OwnerID == userSessionId {
									for index, val := range itemList {
										if val.ID == v.ItemID {
											recordPrint := itemList[index].Name
											recordPrintList = append(recordPrintList, recordPrint)
										}
									}

								}
							}
							fmt.Println(recordPrintList)
							service.AddLog(userSessionId, userSessionRole, "Checked Inventory")
						case 3:
							fmt.Println("--==--==--Item List--==--==--")
							itemList, err := service.GetItems()
							if err != nil {
								log.Fatal(err)
							}
							for index, _ := range itemList {
								fmt.Printf("%d. Item Info : %v\n", index, itemList[index])
								fmt.Println("-------------------------------")
							}
							service.AddLog(userSessionId, userSessionRole, "Listed Items")
						case 4:
							break
						}
					case 3:
						fmt.Println("--==--==--Add Money--==--==--")
						fmt.Println("Plese Type Amount of Money to Add")
						var addMoney int
						fmt.Scanln(&addMoney)
						err := service.AddMoney(userSessionId, addMoney)
						if err != nil {
							fmt.Println(err)
						}
						err = service.AddLog(userSessionId, userSessionRole, "Added Money")
						if err != nil {
							fmt.Println(err)
							return
						}
					case 4:
						return
					}
				}
			} else {
				fmt.Println("User Role Broken")
				return
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
		}
	}
}
