package main

import (
	"fmt"
)

func main() {
	fmt.Println("Were are executing a goroutine")
	arr := []int{2, 3, 4}
	ch := make(chan int, len(arr))

	go func(arr []int, ch chan int) {
		for _, element := range arr {
			ch <- element * 3
		}
	}(arr, ch)

	for i := 0; i < len(arr); i++ {
		fmt.Println("result: %v", <-ch)
	}
}
