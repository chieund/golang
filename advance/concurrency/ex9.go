package main

import "fmt"

func timesThree(arr []int, ch chan int) {
	minusCh := make(chan int, 3)

	for _, element := range arr {
		value := element
		if value%2 == 0 {
			fmt.Println("%2 == 0", value)
			go minusThree(value, minusCh)
			value = <-minusCh
			value = value - 3
		}
		ch <- value
	}

}

func minusThree(number int, ch chan int) {
	ch <- number - 3
	fmt.Println("The functions continues after returning the result")
}

func main() {
	fmt.Println("Were are executing a goroutine")
	arr := []int{6, 9, 12}

	ch := make(chan int, len(arr))
	go timesThree(arr, ch)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("result: %v \n", <-ch)
	}
}
