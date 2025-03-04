package pkg

import (
	"log"
	"os"
	"strconv"
)

func AddLog(userId int, userRole, logMessage string) error {
	logMessage = strconv.Itoa(userId) + " - Identified User " + logMessage + " as a " + userRole + "\n"

	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(logMessage)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
