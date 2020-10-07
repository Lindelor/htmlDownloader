package main

import (
	"io/ioutil"
	"net/http"
)

/*Функция принимает URL, Имя логфайла и папку назначения, записывает тело get запроса в файл*/
func WriteResponse(url, logName, destFilesPath string) {
	response, err := http.Get(url)
	if err != nil {
		writeLog(logName, err.Error())
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			writeLog(logName, err.Error())
		} else {
			writeLog(logName, url+" download is done")
			var filename = urlToName(url)
			output(logName, destFilesPath+"\\"+filename, string(body))
		}
	}
}
