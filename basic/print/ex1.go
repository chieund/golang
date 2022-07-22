package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Print("a", "b", "c")

	fmt.Println("a", "b", "c")

	test := fmt.Sprintf("test %s", "a")

	fmt.Println(test)

	fmt.Printf("test %s \n", "a")

	var num float32 = 7.786
	fmt.Printf("the floating point is %g \n", num)
	fmt.Printf("the floating point is %f \n", num)

	tes2 := fmt.Errorf("show print error %s", "a")
	fmt.Println(tes2)

	fmt.Printf("type %T \n", tes2)

	//log.Fatalln("test")

	log.Panic(tes2)
}
