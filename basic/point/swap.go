package main

import "fmt"
func main() {
	var a int = 100
	var b int = 200

	fmt.Printf("Before value a:%d\n", a)
	fmt.Printf("Before value b:%d\n", b)

	swap(&a, &b)
	fmt.Printf("After value a:%d\n", a)
	fmt.Printf("After value b:%d\n", b)
}

func swap(x *int, y *int) {
	var temp int

	temp = *x
	*x = *y
	*y = temp

}