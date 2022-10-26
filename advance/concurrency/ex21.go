package main

import (
	"fmt"
	"sync"
)

func Hello(ws *sync.WaitGroup) {
	fmt.Println("Hello")
	defer ws.Done()
}
func main() {
	var ws sync.WaitGroup
	ws.Add(2)
	go Hello(&ws)
	ws.Wait()
}
