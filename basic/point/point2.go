package main

import "fmt"

func main() {
	var ptr *int
	var a int = 20

	ptr = &a
	fmt.Printf("The value of ptr is: %x\n", ptr)
	fmt.Printf("value of point: %d", *ptr)
}