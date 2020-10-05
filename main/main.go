package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	/*Определяем флаги*/
	var srcFileName, destFilesPath string
	flag.StringVar(&srcFileName, "filename", "", "URLs file name")
	flag.StringVar(&destFilesPath, "dest", "./", "Dest dir path")
	flag.Parse()

	/*Запускаем таймер*/
	start := time.Now()
	var urlSlice []string

	/*Создаем лог файл*/
	logFileName := "logFile.log"
	logFile, err1 := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer logFile.Close()
	if err1 != nil {
		fmt.Print(fmt.Sprintf("Log file creating is failed. Log in console now\n %s \n", err1.Error()))
	} else {
		log.SetOutput(logFile)
	}

	/*Вызываем функцию, преобразующую файл в слайс урлов*/
	urlSlice, err3 := fileToUrlSlice(srcFileName)
	if err3 != nil {
		writeLog(logFileName, err3.Error())
	} else {
		writeLog(logFileName, fmt.Sprintf("Total strings in source file is: %d", len(urlSlice)))
	}

	/*Отправляем запрос, логируем все действия, создаем файлы с телом ответа*/
	createDirectory(logFileName, destFilesPath)
	for i := 0; i < len(urlSlice); i++ {
		response, err := GetResponse(urlSlice[i])
		if err != nil {
			writeLog(logFileName, err.Error())
		} else {
			writeLog(logFileName, urlSlice[i]+" download is done")
			var filename = urlToName(urlSlice[i])
			output(logFileName, destFilesPath+"\\"+filename, response)
		}
	}

	/*Завершаем замерять время и логируем*/
	duration := time.Since(start)
	workTime := fmt.Sprintf("Work is done by %v", duration)
	writeLog(logFileName, workTime)
}
