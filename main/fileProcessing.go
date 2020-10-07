package main

import (
	"bufio"
	"os"
	"strings"
	"sync"
)

/*Функция для создания директорий, принимает на вход имя логфайла и путь к создаваемой директории*/
func createDirectory(logFileName, filesPath string) {

	err := os.MkdirAll(filesPath, 0777)
	if err != nil {
		writeLog(logFileName, err.Error())
	} else if filesPath == "./" {
		writeLog(logFileName, "Save files in current directory")
	} else {
		writeLog(logFileName, "Directory is created")
	}
}

/*Функция преобразует URL в имя файла для сохранения, принимает URL*/
func urlToName(siteUrl string) string {
	site := strings.ReplaceAll(siteUrl, "http://", "")
	site = strings.ReplaceAll(site, "https://", "")
	site = strings.ReplaceAll(site, ".", "")
	site = strings.ReplaceAll(site, "/", "")
	site += ".html"
	return site
}

/*Функция считывает с файла все строки и преобразует в URL, принимает путь до файла*/
func fileToUrlSlice(filePath string) ([]string, error) {
	var urlSlice = []string{}
	var waitGroup sync.WaitGroup

	file, err := os.Open(filePath)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		waitGroup.Add(1)
		go func(scan string) {
			defer waitGroup.Done()
			urlSlice = append(urlSlice, scan)
		}(scanner.Text())
	}

	waitGroup.Wait()

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	return urlSlice, nil
}
