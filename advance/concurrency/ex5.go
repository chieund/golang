package main

import (
	"fmt"
	"time"
)

func timesThree(number int) {
	fmt.Println(number * 3)
}

func main() {
	fmt.Println("Were are executing a goroutine")
	go timesThree(3)
	fmt.Println("Done!")
	time.Sleep(time.Second)
}
