package main

import (
	"fmt"
)

func numberThree(number int, ch chan int) {
	result := number * 3
	fmt.Println(result)
	ch <- result
}

func main() {
	fmt.Println("Were are executing a goroutine")
	ch := make(chan int)
	go numberThree(3, ch)
	result := <-ch
	fmt.Printf("value result: %v", result)
}
