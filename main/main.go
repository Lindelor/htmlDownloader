package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
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
	var wg sync.WaitGroup

	/*Создаем лог файл*/
	logFileName := "logFile.log"
	logFile, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Print(fmt.Sprintf("Log file creating is failed. Log in console now\n %s \n", err.Error()))
	} else {
		log.SetOutput(logFile)
	}
	if logFile != nil {
		defer logFile.Close()
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
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			WriteResponse(url, logFileName, destFilesPath)
		}(urlSlice[i])
	}

	wg.Wait()

	/*Завершаем замерять время и логируем*/
	duration := time.Since(start)
	workTime := fmt.Sprintf("Work is done by %v", duration)
	writeLog(logFileName, workTime)
}
