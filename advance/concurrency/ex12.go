package main

import "fmt"

func main() {
	//unbuffer := make(chan string)
	//go func(unbuffer chan string) {
	//	unbuffer <- "test"
	//}(unbuffer)
	//fmt.Println(<-unbuffer)

	//unbuffer := make(chan string)
	//go func(){
	//	fmt.Println(<-unbuffer)
	//}()
	//unbuffer <- "test 2"
	//go func(){
	//	fmt.Println(<-unbuffer)
	//}()

	unbuffer := make(chan string)
	go func() {
		unbuffer <- "test"
	}()
	fmt.Println(<-unbuffer)
}
