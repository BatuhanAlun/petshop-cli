package pkg

import (
	"log"
	"os"
)

func AddLog(userId int, logMessage string) error {
	logMessage = logMessage + "\n"
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
