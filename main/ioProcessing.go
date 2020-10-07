package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

/*Функция записи в файл, принимает имя лог файла, имя файла и сообщение для записи*/
func output(logFileName, fileName, message string) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		writeLog(logFileName, err.Error())
	} else {
		_, err := file.WriteString(message + "\n")
		if err != nil {
			writeLog(logFileName, err.Error())
		}
	}
}

/*Функция записи в лог, принимает имя лог файла и сообщение для записи, если файл нельзя открыть -
вывод идет в консоль*/
func writeLog(logName, message string) {

	logfile, err := os.OpenFile(logName, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Print(message + "\n")
	} else {
		log.SetOutput(logfile)
		log.Println(time.Now(), message)
	}
}
