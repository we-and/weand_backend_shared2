package log

import (
	"log"
	"os"
)

type MonylLog struct {
	Warning *log.Logger
	Info    *log.Logger
	Error   *log.Logger
}

var MLogger MonylLog

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	MLogger.Info = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	MLogger.Warning = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	MLogger.Error = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
