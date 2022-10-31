package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func countFirstFile(result chan int, filePath string, keyword string) {
	var numberOffOcc int
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		result <- 0
		return
	}

	numberOffOcc = strings.Count(string(fileContent), keyword)
	result <- numberOffOcc
	defer close(result)
}

func countSecondFile(result chan int, filePath string, keyword string) {
	var numberOffOcc int
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		result <- 0
		return
	}

	numberOffOcc = strings.Count(string(fileContent), keyword)
	result <- numberOffOcc
	defer close(result)
}

func main() {
	countFirstChan := make(chan int)
	countSecondChan := make(chan int)

	go countFirstFile(countFirstChan, "1.txt", "bạn")
	go countSecondFile(countSecondChan, "2.txt", "bạn")

	log.Println("Tổng số lần xuất hiện:", <-countFirstChan+<-countSecondChan)
	log.Println("main finish")
}
