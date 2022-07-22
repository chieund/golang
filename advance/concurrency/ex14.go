package main

import (
	"fmt"
	"sync"
)

func test1(wg *sync.WaitGroup) {
	fmt.Println("Chay ham test 1")
	wg.Done()
}

func test2(wg *sync.WaitGroup) {
	fmt.Println("chay ham test 2")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go test1(&wg)
	go test2(&wg)
	wg.Wait()
	fmt.Println("Good bye")
}
