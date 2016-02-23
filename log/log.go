package log

import (
	"log"
	"os"
)

func init() {
	var err error
	logfile, err = os.OpenFile("pika.log", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln("fail to create pika.log file!")
		os.Exit(-1)
	}
}

var logfile *os.File

func Logger(v ...interface{}) {
	logger := log.New(logfile, "", log.LstdFlags|log.Llongfile)
	logger.Println(v...)
}
