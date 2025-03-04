package pkg

import (
	"log"
	"os"
	"strconv"
)

func AddLog(userId int, userRole, logMessage string) error {
	logMessage = strconv.Itoa(userId) + "- Idied User " + logMessage + "as a " + userRole + "\n"
	file, err := os.OpenFile("log.txt", os.O_WRONLY, 0644) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.WriteString(logMessage)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return nil
}
