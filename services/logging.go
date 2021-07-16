package services

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

const LOGGING = 1
const LOGS_FILE_PATH = "./logs.txt"

func NeedsToShowLogs(command int) bool {
	return command == LOGGING
}

func ShowLogs() {
	fmt.Println("Exibindo Logs")
	file, error := ioutil.ReadFile(LOGS_FILE_PATH)
	if error != nil {
		fmt.Println("Ocorreu um erro", error)
		return
	}
	fmt.Println(string(file))
}

func RegisterLog(site string, status bool) {
	file, error := os.OpenFile(LOGS_FILE_PATH, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if error != nil {
		fmt.Println("Ocorreu um erro", error)
	}
	file.WriteString(time.Now().Format("02/01/2006 15:04:05 ") + site + " - online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}
