package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("Xin chao goroutine")
	fmt.Println("Xin chao main goroutine")
	time.Sleep(time.Second)
}
