package main

import (
	"log"
	"os"
)

func getAccessToken() string {
	data, err := os.ReadFile(".token")
	if err != nil { // если возникла ошибка
		log.Panicln("Unable to create file:", err)
	}
	return string(data)
}
func main() {
	log.Println("start")
	accessToken := getAccessToken()
	log.Println("access token:", accessToken)
	log.Println("finish")
}
