package main

import (
	"fmt"
	"time"
)

func numberThree(number []int, ch chan int) {
	for _, element := range number {
		ch <- element * 3
	}
}

func main() {
	fmt.Println("Were are executing a goroutine")
	var number = []int{1, 2, 3}
	ch := make(chan int)
	go numberThree(number, ch)
	time.Sleep(time.Second)
	for i := 0; i < len(number); i++ {
		result1 := <-ch
		fmt.Println(result1)
	}
	//fmt.Printf("value result: %v", result1)
	//result2 := <-ch
	//result3 := <-ch
	//fmt.Printf("value result: %v, %v, %v", result1, result2, result3)
}
