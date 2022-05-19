package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)

	fmt.Printf("type: %T\n", numbers)
	printSlice(numbers)

	s := make([]int, 0)
	for i := 0; i < 5; i++ {
		s = append(s, i)
		fmt.Printf("\ncap %v, len %v, %p", cap(s), len(s), s)
	}

	slice := []string{99: "", 19: ""}
	fmt.Printf("\nlen slice:%d", len(slice))
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
