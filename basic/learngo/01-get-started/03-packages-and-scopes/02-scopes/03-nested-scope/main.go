package main

import "fmt"
var declareMeAgain = 10

func nested()  {
	var declareMeAgain = 5
	fmt.Println("inside nested:", declareMeAgain)
}

func main() {
	fmt.Println("inside main:", declareMeAgain)

	nested()

	fmt.Println("inside main:", declareMeAgain)
}
