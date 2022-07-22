package main

import (
	"fmt"
	"time"
)

func read(data <-chan string) {
	fmt.Println(<-data)
}

func write(data chan<- string) {
	data <- "test"
}

func main() {
	unbuffer := make(chan string)
	go write(unbuffer)
	go read(unbuffer)

	time.Sleep(time.Second)
}
