package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type lambdalogger struct {
	*log.Logger
	fileName string
}

var (
	FolderPath   string
	LambdaLogger *lambdalogger
	Logger       = CreateLogger()
)

// check if log folder exists
func logFolderExists() {
	FolderPath = "../logs/"
	folderErr := os.MkdirAll(FolderPath, os.ModePerm)
	if folderErr != nil {
		fmt.Println(folderErr)
	}
}

// create lag file if not exist
func createLogFile(filename string) *lambdalogger {
	file, _ := os.OpenFile(FolderPath+filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	return &lambdalogger{
		fileName: filename,
		Logger:   log.New(file, time.Now().Format("2006-01-02 15:04:05")+": Log: ", log.Lshortfile),
	}
}

func CreateLogger() *lambdalogger {
	logFolderExists()
	LambdaLogger = createLogFile("execution_log.log")
	return LambdaLogger
}
