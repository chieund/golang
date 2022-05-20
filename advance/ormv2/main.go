package main

import (
	"gorm-test/controllers"
	"log"
	"os"
)

func main() {
	LOG_FILE := "log/log"

	logWriter, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln("Unable to set logfile:", err.Error())
	}

	log.SetFlags(log.Lshortfile)

	log.SetOutput(logWriter)
	log.Println("This is a log form GoLang")

	controllers.Init()
}
