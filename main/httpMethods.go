package main

import (
	"io/ioutil"
	"net/http"
)

/*Функция принимает URL и возвращает тело get запроса*/
func GetResponse(url string) (string, error) {

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)

	return string(body), nil

}
