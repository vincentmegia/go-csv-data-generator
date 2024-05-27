package main

import (
	"fmt"
	"os"

	"github.com/vincentmegia/go-csv-data-generator/internal/logger"
)

func main() {
	fmt.Println("hello csv data generator")
	// open a file for writing
	// generate fake data for 1m row
	logFile, error := os.Create("log.txt")
	logger := logger.Logger{LogFile: logFile}
	logger.Initialize()
	if error != nil {
		fmt.Println("An error has occured while creating log file: ", error)
		return
	}
	defer logFile.Close()

	file, error := os.Create("user.csv")
	if error != nil {
		fmt.Println("An error has occured: ", error)
		return
	}
	defer file.Close()

	logger.Log(fmt.Sprintf("id,firstname,lastname,username,password,salt\n"))
	for index := 0; index <= 1000000; index++ {
		firstname := fmt.Sprintf("john-%d", index)
		lasttname := fmt.Sprintf("doe-%d", index)
		username := fmt.Sprintf("john.doe-%d", index)
		password := fmt.Sprintf("john.doe.pass-%d", index)
		salt := fmt.Sprintf("salty.john.doe-%d", index)
		row := fmt.Sprintf("%d,%s,%s,%s,%s,%s\n", index, firstname, lasttname, username, password, salt)
		logger.Log("Writing row: ", row)
		file.WriteString(row)
	}
	logger.Log("Exiting gracefully...")
}
