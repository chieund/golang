package main

import "fmt"

func main() {
	var x float64
	x = 20.0

	y := "fff"

	fmt.Println(x)
	fmt.Println(y)

	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)

	var a, b, c = 3, 4, "foo"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Printf("a is type %T\n", a)
	fmt.Printf("b is type %T\n", b)
	fmt.Printf("c is type %T\n", c)
}
